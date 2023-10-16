package dependencies

import (
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/auth"
)

func ResolveAuth() auth.Auth {
	auth := auth.New()

	// Preload a usable user, dont do that in prod :).
	err := auth.SignUp(domain.User{ID: "testUser", Password: "securePassword"})
	if err != nil {
		panic(err)
	}

	return auth
}
