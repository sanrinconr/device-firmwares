package internal

import (
	"time"

	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

type FirmwareRequest struct {
	ID           string    `json:"id"`
	DeviceID     string    `json:"device_id"`
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	ReleaseNotes string    `json:"release_notes"`
	ReleaseDate  time.Time `json:"release_date"`
	URL          string    `json:"url"`
}

func (f FirmwareRequest) ToDomain() (domain.Firmware, error) {
	uuidFirmware, err := uuid.Parse(f.ID)
	if err != nil {
		return domain.Firmware{}, err
	}

	uuidDevice, err := uuid.Parse(f.DeviceID)
	if err != nil {
		return domain.Firmware{}, err
	}

	return domain.Firmware{
		ID:           domain.FirmwareID(uuidFirmware),
		DeviceID:     domain.DeviceID(uuidDevice),
		Name:         f.Name,
		Version:      f.Version,
		ReleaseNotes: f.ReleaseNotes,
		ReleaseDate:  f.ReleaseDate,
		URL:          f.URL,
	}, nil
}
