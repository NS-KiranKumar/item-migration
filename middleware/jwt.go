package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte(os.Getenv("JWT_SECRET"))

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token Misssing", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimSpace(
			strings.TrimPrefix(authHeader, "Bearer "),
		)
		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return secretkey, nil
			},
		)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
