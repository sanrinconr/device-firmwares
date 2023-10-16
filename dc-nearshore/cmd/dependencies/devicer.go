package dependencies

import (
	"github.com/sanrinconr/device-firmwares/dc-nearshore/database"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/devicer"
)

func resolveDevicer(r database.Device) (devicer.Devicer, error) {
	return devicer.New(r)
}
