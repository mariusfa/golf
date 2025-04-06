package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mariusfa/golf/request"
)

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		ctx := r.Context()
		requestIdCtx := ctx.Value(request.RequestIdCtxKey).(*request.RequestIdCtx)
		requestIdCtx.RequestId = requestId

		next.ServeHTTP(w, r)
	})
}
