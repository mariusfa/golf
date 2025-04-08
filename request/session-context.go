package request

import (
	"github.com/mariusfa/golf/auth"
)

type SessionCtx struct {
	Id       string
	Name     string
	Email    string
	Username string
}

func NewSessionCtx(id, name, email, username string) *SessionCtx {
	return &SessionCtx{
		Id:       id,
		Name:     name,
		Email:    email,
		Username: username,
	}
}

func (s *SessionCtx) SetSessionCtx(user auth.AuthUser) {
	s.Id = user.Id
	s.Name = user.Name
	s.Email = user.Email
	s.Username = user.Username
}
