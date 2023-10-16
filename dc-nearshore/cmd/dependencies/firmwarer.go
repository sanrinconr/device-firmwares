package dependencies

import (
	"github.com/sanrinconr/device-firmwares/dc-nearshore/database"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/firmwarer"
)

func resolveFirmwarer(r database.Firmware) (firmwarer.Firmwarer, error) {
	return firmwarer.New(r)
}
