package middleware

import (
	"net/http"
	"strings"

	"github.com/mariusfa/golf/auth"
	"github.com/mariusfa/golf/logging/applog"
	"github.com/mariusfa/golf/request"
)

type AuthParams struct {
	Secret   string
	UserRepo auth.AuthUserRepository
}

func NewAuthParams(secret string, userRepo auth.AuthUserRepository) AuthParams {
	return AuthParams{
		Secret:   secret,
		UserRepo: userRepo,
	}
}

func Auth(next http.Handler, params AuthParams) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			applog.Error("Missing Authorization header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			applog.Error("Invalid Authorization header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := authHeader[len("Bearer "):]
		if token == "" {
			applog.Error("Missing token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId, err := auth.ParseToken(token, params.Secret)
		if err != nil {
			applog.Error("Error parsing token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := params.UserRepo.FindAuthUserById(userId)
		if err != nil {
			applog.Error("Error finding user")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionCtx, ok := r.Context().Value(request.SessionCtxKey).(*request.SessionCtx)
		if !ok {
			applog.Error("Session context not found")
			http.Error(w, "Session context not found", http.StatusInternalServerError)
			return
		}
		sessionCtx.SetSessionCtx(user)

		applog.Info("User authenticated")
		next.ServeHTTP(w, r)
	})
}
