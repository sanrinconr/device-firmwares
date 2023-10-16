package domain

import "fmt"

// ErrInconsistentDomain is generated when the client of the application
// try to insert inconsistent data.
type ErrInconsistentDomain error

// ErrDeviceNotFound is generated when a device not exists.
type ErrDeviceNotFound struct {
	ID DeviceID
}

func (e ErrDeviceNotFound) Error() string {
	return fmt.Sprintf("device '%s' not found", e.ID)
}

// ErrFirmwareFound is generated when a firmware not exists.
type ErrFirmwareNotFound struct {
	ID FirmwareID
}

func (e ErrFirmwareNotFound) Error() string {
	return fmt.Sprintf("firmware '%s' not found", e.ID)
}
