package middleware

import (
	"context"
	"net/http"
	"order-api-start/configs"
	"order-api-start/pkg/jwt"
	"order-api-start/pkg/res"
	"strings"
)

type key string

const contextPhoneKey = "contextPhoneKey"

func Auth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			res.Json(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, phone := jwt.NewJWT(config.Secret).Parse(token)

		if !isValid {
			res.Json(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextPhoneKey, phone)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
