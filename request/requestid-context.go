package request

type RequestIdCtx struct {
	RequestId string
}

func NewRequestIdCtx(requestId string) *RequestIdCtx {
	return &RequestIdCtx{
		RequestId: requestId,
	}
}
