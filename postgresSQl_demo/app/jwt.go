package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"v1/models"
	"v1/utils"

	"github.com/golang-jwt/jwt/v4"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path
		for _, auth := range notAuth {
			if auth == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		response := make(map[string]interface{})
		tokenHandler := r.Header.Get("Auhtorization")
		if tokenHandler == "" {
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}
		spltted := strings.Split(tokenHandler, " ")
		if len(spltted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}
		tokenpart := spltted[1]
		tk := &models.Token{}
		token, err := jwt.ParseWithClaims(tokenpart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil { // Malformed token, returns with http code 403 as usual
			response = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}
		if !token.Valid {
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}
		fmt.Sprintf("User %d", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
