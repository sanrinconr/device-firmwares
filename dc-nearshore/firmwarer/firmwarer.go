package firmwarer

import (
	"context"
	"fmt"

	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
)

const MissingDependency = "missing '%s' in 'firmwarer'"

type Firmwarer struct {
	repository repository
}

type repository interface {
	SelectFirmwareID(context.Context, domain.FirmwareID) (domain.Firmware, error)
	SelectAll(context.Context) (domain.Firmwares, error)
	Insert(context.Context, domain.Firmware, domain.DeviceID) (domain.FirmwareID, error)
	Delete(context.Context, domain.FirmwareID) error
	Update(context.Context, domain.Firmware, domain.DeviceID) error
}

// New creates the implementation of a firmware.
func New(r repository) (Firmwarer, error) {
	if r == nil {
		return Firmwarer{}, fmt.Errorf(MissingDependency, "repository")
	}

	return Firmwarer{
		repository: r,
	}, nil
}

func (f Firmwarer) ObtainByID(ctx context.Context, fID domain.FirmwareID) (domain.Firmware, error) {
	return f.repository.SelectFirmwareID(ctx, fID)
}

func (f Firmwarer) ObtainAll(ctx context.Context) (domain.Firmwares, error) {
	return f.repository.SelectAll(ctx)
}

func (f Firmwarer) Create(
	ctx context.Context,
	firmware domain.Firmware,
	dID domain.DeviceID,
) (domain.FirmwareID, error) {
	return f.repository.Insert(ctx, firmware, dID)
}

func (f Firmwarer) Delete(ctx context.Context, fID domain.FirmwareID) error {
	return f.repository.Delete(ctx, fID)
}

func (f Firmwarer) Update(ctx context.Context, fID domain.Firmware, dID domain.DeviceID) error {
	return f.repository.Update(ctx, fID, dID)
}
