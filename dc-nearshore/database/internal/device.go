package internal

import (
	"time"

	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

type (
	Device struct {
		ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
		Firmwares Firmwares
		Name      string
		CreatedAt time.Time `gorm:"<-:create"`
		UpdatedAt time.Time
	}

	Devices []Device
)

func DeviceFromDomain(d domain.Device) (Device, error) {
	firmwares := make([]Firmware, 0, len(d.Firmwares))
	for _, f := range d.Firmwares {
		firmwares = append(firmwares, FirmwareFromDomain(f, d.ID))
	}

	return Device{
		ID:        uuid.UUID(d.ID),
		Firmwares: firmwares,
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}, nil
}

func (d Devices) ToDomain() domain.Devices {
	res := make(domain.Devices, 0, len(d))
	for _, device := range d {
		res = append(res, device.ToDomain())
	}

	return res
}

func (d Device) ToDomain() domain.Device {
	return domain.Device{
		ID:        domain.DeviceID(d.ID),
		Name:      d.Name,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		Firmwares: d.Firmwares.ToDomain(),
	}
}
