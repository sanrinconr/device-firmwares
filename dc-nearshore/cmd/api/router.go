// Package api setup gin gonic to run, this include:
// - Endpoints creation.
// - Configure middlewares.
// - Obviously, run the gin daemon.
package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/config"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/dependencies"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/middlewares"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/log"
)

// Start create routes, set port and run app.
func Start() {
	router := createRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Printf("%s", err)
	}
}

func createRouter() *gin.Engine {
	err := config.Read()
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Getenv("scopeType"))
	router := gin.New()

	setTrustedProxies(router)

	router.Use(gin.Recovery())
	router.Use(middlewares.LoadLoggerInCtx(logger))
	router.Use(middlewares.Logger)
	router.Use(middlewares.Auth(dependencies.ResolveAuth()))
	router.Use(middlewares.ErrorHandler)
	setRoutes(router)

	return router
}

func setTrustedProxies(r *gin.Engine) {
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}
}
