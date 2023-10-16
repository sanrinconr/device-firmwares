package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/cmd/controller/internal"
)

type Firmware struct {
	firmwarer domain.Firmwarer
}

func NewFirmware(f domain.Firmwarer) (Firmware, error) {
	if f == nil {
		return Firmware{}, fmt.Errorf(MissingDependency, "firmwarer", "controller firmware")
	}

	return Firmware{
		firmwarer: f,
	}, nil
}

func (f Firmware) ObtainFirmwareID(ctx *gin.Context) error {
	id := ctx.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	firmware, err := f.firmwarer.ObtainByID(ctx, domain.FirmwareID(idUUID))
	if err != nil {
		return err
	}

	if firmware.ID == domain.FirmwareID(uuid.Nil) {
		ctx.Status(http.StatusNoContent)

		return nil
	}

	ctx.JSON(http.StatusOK, internal.FirmwareResponseFromDomain(firmware))

	return nil
}

func (f Firmware) ObtainAllFirmwares(ctx *gin.Context) error {
	firmware, err := f.firmwarer.ObtainAll(ctx)
	if err != nil {
		return err
	}

	if len(firmware) == 0 {
		ctx.Status(http.StatusNoContent)

		return nil
	}

	responses := make([]internal.FirmwareResponse, 0, len(firmware))
	for _, firmware := range firmware {
		responses = append(responses, internal.FirmwareResponseFromDomain(firmware))
	}

	ctx.JSON(http.StatusOK, responses)

	return nil
}

func (f Firmware) CreateFirmware(ctx *gin.Context) error {
	deviceReq := new(internal.FirmwareRequest)

	err := ctx.BindJSON(deviceReq)
	if err != nil {
		return err
	}

	deviceUUID, err := uuid.Parse(deviceReq.DeviceID)
	if err != nil {
		return err
	}

	domainDevice, err := deviceReq.ToDomain()
	if err != nil {
		return err
	}

	firmwareID, err := f.firmwarer.Create(ctx, domainDevice, domain.DeviceID(deviceUUID))
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, map[string]string{"created": firmwareID.String()})

	return nil
}

func (f Firmware) UpdateFirmware(ctx *gin.Context) error {
	id := ctx.Param("id")

	firmwareIDUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	firmwareReq := new(internal.FirmwareRequest)

	err = ctx.BindJSON(firmwareReq)
	if err != nil {
		return err
	}

	firmwareReq.ID = firmwareIDUUID.String()

	deviceIDUUID, err := uuid.Parse(firmwareReq.DeviceID)
	if err != nil {
		return err
	}

	domainFirmware, err := firmwareReq.ToDomain()
	if err != nil {
		return err
	}

	err = f.firmwarer.Update(ctx, domainFirmware, domain.DeviceID(deviceIDUUID))
	if err != nil {
		return err
	}

	return nil
}

func (f Firmware) DeleteFirmware(ctx *gin.Context) error {
	id := ctx.Param("id")

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = f.firmwarer.Delete(ctx, domain.FirmwareID(idUUID))
	if err != nil {
		return err
	}

	return nil
}
