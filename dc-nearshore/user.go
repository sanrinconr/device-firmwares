package domain

type User struct {
	ID       string
	Password string
}

type Authenticator interface {
	Valid(User) bool
}
