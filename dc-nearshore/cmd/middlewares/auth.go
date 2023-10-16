package middlewares

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

// Auth ensure request are authenticated.
func Auth(auth domain.Authenticator) func(*gin.Context) {
	return func(ctx *gin.Context) {
		str := ctx.Request.Header.Get("Authorization")

		// Prevent bigger header.
		const bigHeader = 100
		if len(str) > bigHeader {
			ctx.Status(http.StatusForbidden)

			return
		}

		// strip "Basic " from auth.
		authEncoded := strings.ReplaceAll(str, "Basic ", "")

		decodedBytes, err := base64.StdEncoding.DecodeString(authEncoded)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})

			return
		}

		userPass := strings.Split(string(decodedBytes), ":")

		const requiredArgs = 2
		if len(userPass) < requiredArgs {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"msg": "format of auth must be user:password"})

			return
		}

		user := userPass[0]
		pass := userPass[1]

		valid := auth.Valid(domain.User{ID: user, Password: pass})
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusForbidden, map[string]string{"msg": "invalid user/password"})

			return
		}

		ctx.Next()
	}
}
