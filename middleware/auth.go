package middleware

import (
	"net/http"
	"strings"

	"github.com/mariusfa/golf/auth"
	"github.com/mariusfa/golf/request"
)

type AuthParams struct {
	Secret   string
	UserRepo auth.AuthUserRepository
	Logger   loggerPort
}

// TODO: Remove logger from params and use a global logger instead
func NewAuthParams(secret string, userRepo auth.AuthUserRepository, logger loggerPort) AuthParams {
	return AuthParams{
		Secret:   secret,
		UserRepo: userRepo,
		Logger:   logger,
	}
}

func Auth(next http.Handler, params AuthParams) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-Id")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			params.Logger.Error("Missing Authorization header", requestId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			params.Logger.Error("Invalid Authorization header", requestId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := authHeader[len("Bearer "):]
		if token == "" {
			params.Logger.Error("Missing token", requestId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId, err := auth.ParseToken(token, params.Secret)
		if err != nil {
			params.Logger.Error("Error parsing token", requestId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := params.UserRepo.FindAuthUserById(userId)
		if err != nil {
			params.Logger.Error("Error finding user", requestId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionCtx, ok := r.Context().Value(request.SessionCtxKey).(*request.SessionCtx)
		if !ok {
			params.Logger.Error("Session context not found", requestId)
			http.Error(w, "Session context not found", http.StatusInternalServerError)
			return
		}
		sessionCtx.SetSessionCtx(user)

		params.Logger.Info("User authenticated", requestId)
		next.ServeHTTP(w, r)
	})
}

type loggerPort interface {
	Info(message string, requestId string)
	Error(message string, requestId string)
}
