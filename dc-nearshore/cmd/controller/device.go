package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/controller/internal"
)

const MissingDependency = "missing '%s' in '%s'"

type Device struct {
	devicer domain.Devicer
}

func NewDevice(d domain.Devicer) (Device, error) {
	if d == nil {
		return Device{}, fmt.Errorf(MissingDependency, "devicer", "controller device")
	}

	return Device{
		devicer: d,
	}, nil
}

func (d Device) ObtainDeviceID(ctx *gin.Context) error {
	id := ctx.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	device, err := d.devicer.ObtainByID(ctx, domain.DeviceID(idUUID))
	if err != nil {
		return err
	}

	if device.ID == domain.DeviceID(uuid.Nil) {
		ctx.Status(http.StatusNoContent)

		return nil
	}

	ctx.JSON(http.StatusOK, internal.DeviceReponseFromDomain(device))

	return nil
}

func (d Device) ObtainAllDevices(ctx *gin.Context) error {
	devices, err := d.devicer.ObtainAll(ctx)
	if err != nil {
		return err
	}

	if len(devices) == 0 {
		ctx.Status(http.StatusNoContent)

		return nil
	}

	responses := make([]internal.DeviceResponse, 0, len(devices))
	for _, device := range devices {
		responses = append(responses, internal.DeviceReponseFromDomain(device))
	}

	ctx.JSON(http.StatusOK, responses)

	return nil
}

func (d Device) CreateDevice(ctx *gin.Context) error {
	deviceReq := new(internal.DeviceRequest)

	err := ctx.BindJSON(deviceReq)
	if err != nil {
		return err
	}

	deviceDomain, err := deviceReq.ToDomain()
	if err != nil {
		return err
	}

	deviceID, err := d.devicer.Create(ctx, deviceDomain)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, map[string]string{"created": deviceID.String()})

	return nil
}

func (d Device) UpdateDevice(ctx *gin.Context) error {
	id := ctx.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	deviceReq := new(internal.DeviceRequest)
	deviceReq.ID = idUUID.String()

	err = ctx.BindJSON(deviceReq)
	if err != nil {
		return err
	}

	deviceDomain, err := deviceReq.ToDomain()
	if err != nil {
		return err
	}

	err = d.devicer.Update(ctx, deviceDomain)
	if err != nil {
		return err
	}

	return nil
}

func (d Device) DeleteDevice(ctx *gin.Context) error {
	id := ctx.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = d.devicer.Delete(ctx, domain.DeviceID(idUUID))
	if err != nil {
		return err
	}

	return nil
}
