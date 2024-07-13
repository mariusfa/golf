package middleware

import (
	"context"
	"net/http"

	"github.com/mariusfa/golf/auth"
)

type AuthParams struct {
	Secret   string
	UserRepo userRepositoryPort
	Logger   loggerPort
}

func NewAuthParams(secret string, userRepo userRepositoryPort, logger loggerPort) AuthParams {
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

		user, err := params.UserRepo.FindById(userId)
		if err != nil {
			params.Logger.Error("Error finding user", requestId)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		contextUser := r.Context().Value(UserKey)
		if contextUser != nil {
			contextUser = user
		}

		params.Logger.Info("User authenticated", requestId)
		ctx := context.WithValue(r.Context(), UserKey, contextUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type loggerPort interface {
	Info(message string, requestId string)
	Error(message string, requestId string)
}

type userRepositoryPort interface {
	FindById(userId string) (*User, error)
}

type User struct {
	Id   string
	Name string
}
