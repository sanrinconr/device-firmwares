package internal_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/database/internal"
	"github.com/stretchr/testify/assert"
)

var (
	defaultTime = time.Date(2022, 10, 10, 10, 10, 10, 0, time.UTC)
	defaultUUID = uuid.New()
)

func TestDeviceFromDomain_Success(t *testing.T) {
	d := domain.Device{
		ID:        domain.DeviceID(defaultUUID),
		Name:      "testDevice",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
		Firmwares: []domain.Firmware{
			{
				ID:           domain.FirmwareID(defaultUUID),
				Name:         "testFirmware",
				Version:      "1.0.0",
				ReleaseNotes: "some release",
				ReleaseDate:  defaultTime,
				URL:          "https://test.com",
				CreatedAt:    defaultTime,
				UpdatedAt:    defaultTime,
			},
		},
	}
	want := defaultDevice()

	got, err := internal.DeviceFromDomain(d)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestDevicesToDomain_Success(t *testing.T) {
	devices := internal.Devices{
		defaultDevice(),
		defaultDevice(),
	}
	want := domain.Device{
		ID:        domain.DeviceID(defaultUUID),
		Name:      "testDevice",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
		Firmwares: []domain.Firmware{
			{
				ID:           domain.FirmwareID(defaultUUID),
				DeviceID:     domain.DeviceID(defaultUUID),
				Name:         "testFirmware",
				Version:      "1.0.0",
				ReleaseNotes: "some release",
				ReleaseDate:  defaultTime,
				URL:          "https://test.com",
				CreatedAt:    defaultTime,
				UpdatedAt:    defaultTime,
			},
		},
	}

	got := devices.ToDomain()[1]

	assert.Equal(t, want, got)
}

func defaultDevice() internal.Device {
	return internal.Device{
		ID:        defaultUUID,
		Name:      "testDevice",
		CreatedAt: defaultTime,
		UpdatedAt: defaultTime,
		Firmwares: []internal.Firmware{
			{
				ID:           defaultUUID,
				DeviceID:     defaultUUID,
				Name:         "testFirmware",
				Version:      "1.0.0",
				ReleaseNotes: "some release",
				ReleaseDate:  defaultTime,
				URL:          "https://test.com",
				CreatedAt:    defaultTime,
				UpdatedAt:    defaultTime,
			},
		},
	}
}
