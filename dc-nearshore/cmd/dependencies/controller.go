package dependencies

import (
	"github.com/gin-gonic/gin"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/controller"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/devicer"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/firmwarer"
)

type Controller struct {
	Ping               func(*gin.Context) error
	ObtainDeviceID     func(*gin.Context) error
	ObtainAllDevices   func(*gin.Context) error
	CreateDevice       func(*gin.Context) error
	UpdateDevice       func(*gin.Context) error
	DeleteDevice       func(*gin.Context) error
	ObtainFirmwareID   func(*gin.Context) error
	ObtainAllFirmwares func(*gin.Context) error
	CreateFirmware     func(*gin.Context) error
	UpdateFirmware     func(*gin.Context) error
	DeleteFirmware     func(*gin.Context) error
}

func ResolveControllers() Controller {
	packages := ResolvePackages()

	device := resolveDeviceControllers(packages.Devicer)
	firmware := resolveFirmwareControllers(packages.Firmwarer)

	return Controller{
		Ping:               controller.NewPing(),
		ObtainDeviceID:     device.ObtainDeviceID,
		ObtainAllDevices:   device.ObtainAllDevices,
		CreateDevice:       device.CreateDevice,
		UpdateDevice:       device.UpdateDevice,
		DeleteDevice:       device.DeleteDevice,
		ObtainFirmwareID:   firmware.ObtainFirmwareID,
		ObtainAllFirmwares: firmware.ObtainAllFirmwares,
		CreateFirmware:     firmware.CreateFirmware,
		UpdateFirmware:     firmware.UpdateFirmware,
		DeleteFirmware:     firmware.DeleteFirmware,
	}
}

func resolveDeviceControllers(d devicer.Devicer) controller.Device {
	device, err := controller.NewDevice(d)
	if err != nil {
		panic(err)
	}

	return device
}

func resolveFirmwareControllers(f firmwarer.Firmwarer) controller.Firmware {
	firmware, err := controller.NewFirmware(f)
	if err != nil {
		panic(err)
	}

	return firmware
}
