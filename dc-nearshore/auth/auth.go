package auth

import (
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
	"golang.org/x/crypto/bcrypt"
)

// Auth is a dummy implementation of authentication in memory.
type Auth struct {
	db map[string][]byte
}

func New() Auth {
	return Auth{
		db: map[string][]byte{},
	}
}

func (a Auth) Valid(u domain.User) bool {
	if err := bcrypt.CompareHashAndPassword(a.db[u.ID], []byte(u.Password)); err != nil {
		return false
	}

	return true
}

func (a Auth) SignUp(u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.db[u.ID] = hash

	return nil
}
