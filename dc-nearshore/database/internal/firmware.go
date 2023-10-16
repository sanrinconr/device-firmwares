package internal

import (
	"time"

	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

type (
	Firmware struct {
		ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
		DeviceID     uuid.UUID
		Name         string
		Version      string
		ReleaseNotes string
		ReleaseDate  time.Time
		URL          string
		CreatedAt    time.Time `gorm:"<-:create"`
		UpdatedAt    time.Time
	}

	Firmwares []Firmware
)

func (f Firmwares) ToDomain() []domain.Firmware {
	res := make([]domain.Firmware, 0, len(f))
	for _, firmware := range f {
		res = append(res, firmware.ToDomain())
	}

	return res
}

func (f Firmware) ToDomain() domain.Firmware {
	return domain.Firmware{
		ID:           domain.FirmwareID(f.ID),
		DeviceID:     domain.DeviceID(f.DeviceID),
		Name:         f.Name,
		Version:      f.Version,
		ReleaseNotes: f.ReleaseNotes,
		ReleaseDate:  f.ReleaseDate,
		URL:          f.URL,
		CreatedAt:    f.CreatedAt,
		UpdatedAt:    f.UpdatedAt,
	}
}

func FirmwareFromDomain(f domain.Firmware, dID domain.DeviceID) Firmware {
	return Firmware{
		ID:           uuid.UUID(f.ID),
		DeviceID:     uuid.UUID(dID),
		Name:         f.Name,
		Version:      f.Version,
		ReleaseNotes: f.ReleaseNotes,
		ReleaseDate:  f.ReleaseDate,
		URL:          f.URL,
		CreatedAt:    f.CreatedAt,
		UpdatedAt:    f.UpdatedAt,
	}
}
