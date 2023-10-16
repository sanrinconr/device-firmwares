package internal

import (
	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

type DeviceRequest struct {
	ID        string            `json:"id,omitempty"`
	Name      string            `json:"name"`
	Firmwares []FirmwareRequest `json:"firmware,omitempty"`
}

func (d DeviceRequest) ToDomain() (domain.Device, error) {
	firmwares := make(domain.Firmwares, 0, len(d.Firmwares))

	for _, firmware := range d.Firmwares {
		firmwares = append(firmwares, domain.Firmware{
			Name:         firmware.Name,
			Version:      firmware.Version,
			ReleaseNotes: firmware.ReleaseNotes,
			ReleaseDate:  firmware.ReleaseDate,
			URL:          firmware.URL,
		})
	}

	uuidDevice, err := uuid.Parse(d.ID)
	if uuidDevice != uuid.Nil && err != nil {
		return domain.Device{}, err
	}

	return domain.Device{
		ID:        domain.DeviceID(uuidDevice),
		Name:      d.Name,
		Firmwares: firmwares,
	}, nil
}
