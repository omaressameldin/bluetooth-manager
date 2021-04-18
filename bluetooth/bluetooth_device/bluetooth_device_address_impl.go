package bluetooth_device

import (
	"fmt"
	"regexp"
)

var correctFormat = "[HEX][HEX]:[HEX][HEX]:[HEX][HEX]:[HEX][HEX]:[HEX][HEX]:[HEX][HEX]"
var ErrDeviceAddress = fmt.Errorf("wrong format expected [%s]", correctFormat)

type bluetoothDeviceAddressImpl struct {
	address string
}

func CreateBluetoothDeviceAddress(address string) (BluetoothDeviceAddress, error) {
	err := validateAddress(address)
	if err != nil {
		return nil, err
	}

	return &bluetoothDeviceAddressImpl{
		address: address,
	}, nil
}

func validateAddress(address string) error {
	r, err := regexp.Compile("([0-9A-F]{2}):([0-9A-F]{2}):([0-9A-F]{2}):([0-9A-F]{2}):([0-9A-F]{2}):([0-9A-F]{2})")
	if err != nil {
		return err
	}

	if !r.MatchString(address) {
		return fmt.Errorf("%w got %s", ErrDeviceAddress, address)
	}

	return nil
}

func (bda bluetoothDeviceAddressImpl) String() string {
	return bda.address
}
