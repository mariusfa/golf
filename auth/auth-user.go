package auth

type AuthUser struct {
	Id             string
	Name           string
	HashedPassword string
	Email          string
	Username       string
}

type AuthUserRepository interface {
	FindAuthUserById(userId string) (AuthUser, error)
}
