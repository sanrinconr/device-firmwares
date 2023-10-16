package internal

import (
	"time"

	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

type DeviceResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	FirmwareIDs []string  `json:"firmware_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func DeviceReponseFromDomain(d domain.Device) DeviceResponse {
	firmwaresIDs := make([]string, 0, len(d.Firmwares))

	for _, firmware := range d.Firmwares {
		firmwaresIDs = append(firmwaresIDs, firmware.ID.String())
	}

	return DeviceResponse{
		ID:          d.ID.String(),
		Name:        d.Name,
		FirmwareIDs: firmwaresIDs,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
