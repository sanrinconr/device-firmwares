package devicer

import (
	"context"
	"fmt"

	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

const MissingDependency = "missing '%s' in 'devicer'"

type Devicer struct {
	repository repository
}

type repository interface {
	SelectDeviceID(ctx context.Context, id domain.DeviceID) (domain.Device, error)
	SelectAll(ctx context.Context) (domain.Devices, error)
	Insert(ctx context.Context, device domain.Device) (domain.DeviceID, error)
	Delete(ctx context.Context, deviceID domain.DeviceID) error
	Update(ctx context.Context, device domain.Device) error
}

// New creates the implementation of a device.
func New(r repository) (Devicer, error) {
	if r == nil {
		return Devicer{}, fmt.Errorf(MissingDependency, "repository")
	}

	return Devicer{
		repository: r,
	}, nil
}

func (d Devicer) ObtainByID(ctx context.Context, dID domain.DeviceID) (domain.Device, error) {
	return d.repository.SelectDeviceID(ctx, dID)
}

func (d Devicer) ObtainAll(ctx context.Context) (domain.Devices, error) {
	return d.repository.SelectAll(ctx)
}

func (d Devicer) Create(ctx context.Context, device domain.Device) (domain.DeviceID, error) {
	return d.repository.Insert(ctx, device)
}

func (d Devicer) Delete(ctx context.Context, dID domain.DeviceID) error {
	return d.repository.Delete(ctx, dID)
}

func (d Devicer) Update(ctx context.Context, device domain.Device) error {
	return d.repository.Update(ctx, device)
}
