package database

import (
	"time"

	"github.com/sanrinconr/device-firmwares/dc-nearshore/database/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Device
	Firmware
}

func New(dsn string) (Database, error) {
	gormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: dsn,
			},
		),
		&gorm.Config{
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		},
	)
	if err != nil {
		return Database{}, err
	}

	err = gormDB.AutoMigrate(
		internal.Device{},
		internal.Firmware{},
	)
	if err != nil {
		return Database{}, err
	}

	return Database{
		Device:   Device{driver: gormDB},
		Firmware: Firmware{driver: gormDB},
	}, nil
}
