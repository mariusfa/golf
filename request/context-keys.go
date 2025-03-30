package request

type CtxKey string

const (
	SessionCtxKey   CtxKey = CtxKey("sessionCtx")
	RequestIdCtxKey CtxKey = CtxKey("requestIdCtx")
)
