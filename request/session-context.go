package request

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
