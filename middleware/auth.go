package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("JWT_SECRET")
var jwtSecret = []byte(secret)

func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		//"role":   role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, "")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimSpace(parts[1])
		//parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid Token %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "unexpected token", http.StatusUnauthorized)
			return
		}
		// extract user Id from claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "unable to extract the claims", http.StatusInternalServerError)
			return
		}
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "unable to extract user", http.StatusUnauthorized)
			return
		}
		userID := uint(userIDFloat)
		// add user id to context
		ctx := context.WithValue(r.Context(), "user_id", userID)
		// moving to next protected endpoint
		next(w, r.WithContext(ctx))
	}
}
