package request

import "context"

type RequestIdCtx struct {
	RequestId string
}

func NewRequestIdCtx(requestId string) *RequestIdCtx {
	return &RequestIdCtx{
		RequestId: requestId,
	}
}

func WithRequestIdCtx(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, RequestIdCtxKey, NewRequestIdCtx(requestId))
}
