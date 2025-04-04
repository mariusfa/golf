package request

import "context"

type SessionCtx struct {
	Username string
	UserId   string
}

func NewSessionCtx(username, userId string) *SessionCtx {
	return &SessionCtx{
		Username: username,
		UserId:   userId,
	}
}

func WithSessionCtx(ctx context.Context, username, userId string) context.Context {
	return context.WithValue(ctx, SessionCtxKey, NewSessionCtx(username, userId))
}

/*func GetSessionCtx(ctx context.Context) *SessionCtx {
	if ctx == nil {
		return nil
	}
	if sessionCtx, ok := ctx.Value(SessionCtxKey).(*SessionCtx); ok {
		return sessionCtx
	}
	return nil
}
*/
