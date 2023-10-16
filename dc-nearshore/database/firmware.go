package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	domain "github.com/sanrinconr/device-firmwares/dc-nearshore"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/database/internal"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/log"
	"gorm.io/gorm"
)

type Firmware struct {
	driver *gorm.DB
}

func (f Firmware) SelectFirmwareID(ctx context.Context, id domain.FirmwareID) (domain.Firmware, error) {
	firmware := new(internal.Firmware)

	res := f.driver.WithContext(ctx).First(&firmware, "id = ?", id)
	if err := res.Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return domain.Firmware{}, nil
		}

		return domain.Firmware{}, err
	}

	return firmware.ToDomain(), nil
}

func (f Firmware) SelectAll(ctx context.Context) (domain.Firmwares, error) {
	firmwares := make(internal.Firmwares, 0)

	res := f.driver.WithContext(ctx).Find(&firmwares)
	if err := res.Error; err != nil {
		return nil, err
	}

	return firmwares.ToDomain(), nil
}

func (f Firmware) Insert(
	ctx context.Context,
	firmware domain.Firmware,
	dID domain.DeviceID,
) (domain.FirmwareID, error) {
	entity := internal.FirmwareFromDomain(firmware, dID)

	res := f.driver.Create(&entity)
	if err := res.Error; err != nil {
		return domain.FirmwareID{}, err
	}

	return domain.FirmwareID(entity.ID), nil
}

func (f Firmware) Delete(ctx context.Context, firmwareID domain.FirmwareID) error {
	return f.driver.Delete(&internal.Firmware{}, "id = ?", firmwareID.String()).Error
}

func (f Firmware) Update(ctx context.Context, firmware domain.Firmware, dID domain.DeviceID) error {
	entity := internal.FirmwareFromDomain(firmware, dID)

	exists, err := f.exists(ctx, firmware.ID)
	if err != nil {
		return err
	}

	if !exists {
		log.Info(ctx, fmt.Sprintf("ignoring update, device id dont exists '%s'", firmware.ID.String()))

		return nil
	}

	res := f.driver.Save(&entity)

	if res.RowsAffected == 0 {
		log.Info(ctx, fmt.Sprintf("no firmware updated id '%s' not exists", firmware.ID))

		return domain.ErrFirmwareNotFound{
			ID: firmware.ID,
		}
	}

	return nil
}

func (f Firmware) exists(ctx context.Context, dID domain.FirmwareID) (bool, error) {
	findedFirmware, err := f.SelectFirmwareID(ctx, dID)
	if err != nil {
		return false, err
	}

	if findedFirmware.ID == domain.FirmwareID(uuid.Nil) {
		return false, nil
	}

	return true, nil
}
