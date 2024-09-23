package middleware

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
)

const AuthUserID string = "middleware.auth.userID"

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		// Check that header begins with a prefix of Bearer
		if !strings.HasPrefix(authorization, "Bearer ") {
			writeUnauthed(w)
			return
		}

		// get the token from request header
		encodedToken := strings.TrimPrefix(authorization, "Bearer ")

		// Decode the token from base64
		token, err := base64.StdEncoding.DecodeString(encodedToken)
		if err != nil {
			writeUnauthed(w)
			return
		}

		// Assume the base64 token is valid user id.
		userID := string(token)
		// fmt.Println("user ID: ", userID)
		ctx := context.WithValue(r.Context(), AuthUserID, userID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
