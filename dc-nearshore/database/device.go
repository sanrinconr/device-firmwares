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

type Device struct {
	driver *gorm.DB
}

func (d Device) SelectDeviceID(ctx context.Context, id domain.DeviceID) (domain.Device, error) {
	device := new(internal.Device)

	res := d.driver.WithContext(ctx).Preload("Firmwares").First(&device, "id = ?", id.String())
	if err := res.Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return domain.Device{}, nil
		}

		return domain.Device{}, err
	}

	return device.ToDomain(), nil
}

func (d Device) SelectAll(ctx context.Context) (domain.Devices, error) {
	devices := make(internal.Devices, 0)

	res := d.driver.WithContext(ctx).Preload("Firmwares").Find(&devices)
	if err := res.Error; err != nil {
		return nil, err
	}

	return devices.ToDomain(), nil
}

func (d Device) Insert(ctx context.Context, device domain.Device) (domain.DeviceID, error) {
	entity, err := internal.DeviceFromDomain(device)
	if err != nil {
		return domain.DeviceID{}, err
	}

	res := d.driver.Create(&entity)
	if err := res.Error; err != nil {
		return domain.DeviceID{}, err
	}

	return domain.DeviceID(entity.ID), nil
}

func (d Device) Delete(ctx context.Context, deviceID domain.DeviceID) error {
	err := d.driver.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&internal.Firmware{}, "device_id = ?", deviceID.String())
		if err := res.Error; err != nil {
			return err
		}

		res = tx.Delete(&internal.Device{}, "id = ?", deviceID.String())
		if err := res.Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func (d Device) Update(ctx context.Context, device domain.Device) error {
	entity, err := internal.DeviceFromDomain(device)
	if err != nil {
		return err
	}

	exists, err := d.exists(ctx, device.ID)
	if err != nil {
		return err
	}

	if !exists {
		log.Info(ctx, fmt.Sprintf("ignoring update, device id dont exists '%s'", device.ID.String()))

		return nil
	}

	res := d.driver.Save(&entity)

	if res.RowsAffected == 0 {
		log.Info(ctx, fmt.Sprintf("no device updated id '%s' not exists", device.ID))

		return domain.ErrDeviceNotFound{
			ID: device.ID,
		}
	}

	return nil
}

func (d Device) exists(ctx context.Context, dID domain.DeviceID) (bool, error) {
	findedDevice, err := d.SelectDeviceID(ctx, dID)
	if err != nil {
		return false, err
	}

	if findedDevice.ID == domain.DeviceID(uuid.Nil) {
		return false, nil
	}

	return true, nil
}
