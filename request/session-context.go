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

func (s *SessionCtx) SetSessionCtx(auth.AuthUser) {
	s.Id = s.Id
	s.Name = s.Name
	s.Email = s.Email
	s.Username = s.Username
}
