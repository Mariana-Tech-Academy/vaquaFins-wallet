package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var users = map[string]User{
	"1": {ID: "1", Username: "user1", Password: "password1", Role: "user"},
}

var jwtSecret = []byte("your-secret-key")

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func GenerateJWT(userID, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var foundUser User
	for _, user := range users {
		if user.Username == username {
			foundUser = user
			break
		}
	}

	if foundUser.ID == "" {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !foundUser.CheckPassword(password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(foundUser.ID, foundUser.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
		
		// Log the request details
		start := time.Now()
        fmt.Printf("Request: %s %s\n at %s\n" , r.Method, r.URL.Path, start.Format(time.RFC1123))

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}
func AuthMiddleware(next http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, "")
		if len(parts) != 2 || parts[0] != "Bearer"{
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		_,err := VerifyJWT(tokenString)
		if err != nil{
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next(w, r)

}}
