package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type BlacklistChecker interface {
	IsTokenBlacklisted(jti string) (bool, error)
}

var secret = os.Getenv("JWT_SECRET")
var jwtSecret = []byte(secret)

func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"jti":     fmt.Sprintf("%d-%d", userID, time.Now().UnixNano()), // Added JTI claim
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

/*func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
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
}*/

func AuthMiddleware(repo BlacklistChecker) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "missing token", http.StatusUnauthorized)
                return
            }
            parts := strings.Split(authHeader, " ")
            if len(parts) != 2 || parts[0] != "Bearer" {
                http.Error(w, "invalid token format", http.StatusUnauthorized)
                return
            }
            tokenString := strings.TrimSpace(parts[1])

            // Parse and validate the token
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

            // Extract claims to get the JTI for the blacklist check
            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok {
                http.Error(w, "unable to extract the claims", http.StatusInternalServerError)
                return
            }

            jti, ok := claims["jti"].(string)
            if !ok {
                http.Error(w, "JTI claim not found in token", http.StatusUnauthorized)
                return
            }

            // The Blacklist Check: Call the repository to see if the JTI exists
            isBlacklisted, err := repo.IsTokenBlacklisted(jti)
            if err != nil && err != gorm.ErrRecordNotFound {
                http.Error(w, "error checking token validity", http.StatusInternalServerError)
                return
            }
            if isBlacklisted {
                http.Error(w, "token has been logged out", http.StatusUnauthorized)
                return
            }

            
            userIDFloat, ok := claims["user_id"].(float64)
            if !ok {
                http.Error(w, "unable to extract user", http.StatusUnauthorized)
                return
            }
            userID := uint(userIDFloat)
            ctx := context.WithValue(r.Context(), "user_id", userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}


/*import (
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
}*/
