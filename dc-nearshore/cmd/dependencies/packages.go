package dependencies

import (
	"github.com/sanrinconr/device-firmwares/dc-nearshore/devicer"
	"github.com/sanrinconr/device-firmwares/dc-nearshore/firmwarer"
)

type Packages struct {
	Devicer   devicer.Devicer
	Firmwarer firmwarer.Firmwarer
}

func ResolvePackages() Packages {
	db, err := resolveDatabase()
	if err != nil {
		panic(err)
	}

	devicer, err := resolveDevicer(db.Device)
	if err != nil {
		panic(err)
	}

	firmwarer, err := resolveFirmwarer(db.Firmware)
	if err != nil {
		panic(err)
	}

	return Packages{
		Devicer:   devicer,
		Firmwarer: firmwarer,
	}
}
