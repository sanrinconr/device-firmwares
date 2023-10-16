package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	// Firmware is the loaded firmware to a device.
	Firmware struct {
		ID           FirmwareID
		DeviceID     DeviceID
		Name         string
		Version      string
		ReleaseNotes string
		ReleaseDate  time.Time
		URL          string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	FirmwareID uuid.UUID

	Firmwares []Firmware
)

// Firmwarer represent the behavior of a firmware.
type Firmwarer interface {
	ObtainByID(context.Context, FirmwareID) (Firmware, error)
	ObtainAll(context.Context) (Firmwares, error)
	Create(context.Context, Firmware, DeviceID) (FirmwareID, error)
	Delete(context.Context, FirmwareID) error
	Update(context.Context, Firmware, DeviceID) error
}

func (f FirmwareID) String() string {
	parsed := uuid.UUID(f)

	return parsed.String()
}
