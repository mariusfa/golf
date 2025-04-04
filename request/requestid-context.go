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

/*func GetRequestIdCtx(ctx context.Context) *RequestIdCtx {
	if ctx == nil {
		return nil
	}
	if requestIdCtx, ok := ctx.Value(RequestIdCtxKey).(*RequestIdCtx); ok {
		return requestIdCtx
	}
	return nil
}
*/
