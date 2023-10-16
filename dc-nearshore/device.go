// Package domain have the bussines artefacts that are transversal to the bussines.
package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	// Device where a firmware is loaded.
	Device struct {
		ID        DeviceID
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Firmwares []Firmware
	}

	DeviceID uuid.UUID
	Devices  []Device
)

// Devicer define the behavior of a device.
type Devicer interface {
	ObtainByID(context.Context, DeviceID) (Device, error)
	ObtainAll(context.Context) (Devices, error)
	Create(context.Context, Device) (DeviceID, error)
	Delete(context.Context, DeviceID) error
	Update(context.Context, Device) error
}

func (d DeviceID) String() string {
	parsed := uuid.UUID(d)

	return parsed.String()
}
