package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Anwesa-s/AKSH/internal/auth"
	"github.com/golang-jwt/jwt/v5"
)

type authContextKey string

const (
	userIDKey authContextKey = "userID"
	roleKey   authContextKey = "role"
)

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Login does not require authentication.
		if r.URL.Path == "/login" || r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := auth.ValidateToken(tokenString)

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID := claims["user_id"]
		role := claims["role"]

		ctx := context.WithValue(
			r.Context(),
			userIDKey,
			userID,
		)

		ctx = context.WithValue(
			ctx,
			roleKey,
			role,
		)

		r = r.WithContext(ctx)

		fmt.Printf(
			"Authenticated user_id=%v role=%v\n",
			userID,
			role,
		)

		next.ServeHTTP(w, r)
	})
}
func GetUserID(r *http.Request) interface{} {
	return r.Context().Value(userIDKey)
}

func GetRole(r *http.Request) interface{} {
	return r.Context().Value(roleKey)
}
