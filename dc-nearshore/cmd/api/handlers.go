package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/dependencies"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/middlewares"
)

func setRoutes(router *gin.Engine) {
	resolver := dependencies.ResolveControllers()

	public := router.Group("")
	public.GET("/ping", middlewares.Bridge(resolver.Ping))

	public.GET("/devices/:id", middlewares.Bridge(resolver.ObtainDeviceID))
	public.GET("/devices", middlewares.Bridge(resolver.ObtainAllDevices))
	public.POST("/devices", middlewares.Bridge(resolver.CreateDevice))
	public.PUT("/devices/:id", middlewares.Bridge(resolver.UpdateDevice))
	public.DELETE("/devices/:id", middlewares.Bridge(resolver.DeleteDevice))

	public.GET("/firmwares/:id", middlewares.Bridge(resolver.ObtainFirmwareID))
	public.GET("/firmwares", middlewares.Bridge(resolver.ObtainAllFirmwares))
	public.POST("/firmwares", middlewares.Bridge(resolver.CreateFirmware))
	public.PUT("/firmwares/:id", middlewares.Bridge(resolver.UpdateFirmware))
	public.DELETE("/firmwares/:id", middlewares.Bridge(resolver.DeleteFirmware))
}
