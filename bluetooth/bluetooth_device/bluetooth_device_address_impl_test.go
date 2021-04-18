package bluetooth_device

import (
	"errors"
	"testing"
)

type createAddressTestTable struct {
	address        string
	expectedOutput error
}

func TestCreateAddress(t *testing.T) {
	testTable := []createAddressTestTable{
		{
			address:        "invalid",
			expectedOutput: ErrDeviceAddress,
		},
		{
			address:        "11:22:33:44:55:66",
			expectedOutput: nil,
		},
	}

	for _, row := range testTable {
		address, err := CreateBluetoothDeviceAddress(row.address)
		if !errors.Is(err, row.expectedOutput) {
			t.Errorf("Wrong error type expected type %T, got %T", row.expectedOutput, err)
		}
		if address != nil && address.String() != row.address {
			t.Errorf("wrong address epected %s , got %s", row.address, address)
		}
	}
}
