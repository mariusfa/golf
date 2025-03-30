package utils

import (
	"context"

	"github.com/mariusfa/golf/request"
)

func ExtractFromContext(ctx context.Context) (username string, requestId string) {
	sessionCtx, ok := ctx.Value(request.SessionCtxKey).(*request.SessionCtx)
	if ok {
		username = sessionCtx.Username
	}

	requestIdCtx, ok := ctx.Value(request.RequestIdCtxKey).(*request.RequestIdCtx)
	if ok {
		requestId = requestIdCtx.RequestId
	}

	return
}
