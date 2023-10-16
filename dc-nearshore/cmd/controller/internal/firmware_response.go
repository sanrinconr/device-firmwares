package internal

import (
	"time"

	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

type FirmwareResponse struct {
	ID           string    `json:"id"`
	DeviceID     string    `json:"device_id"`
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	ReleaseNotes string    `json:"release_notes"`
	ReleaseDate  time.Time `json:"release_date"`
	URL          string    `json:"url,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FirmwareResponseFromDomain(d domain.Firmware) FirmwareResponse {
	return FirmwareResponse{
		ID:           d.ID.String(),
		DeviceID:     d.DeviceID.String(),
		Name:         d.Name,
		Version:      d.Version,
		ReleaseNotes: d.ReleaseNotes,
		ReleaseDate:  d.ReleaseDate,
		URL:          d.URL,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}
