package bluetooth_device

import (
	"errors"
	"testing"
)

type createDeviceTestTable struct {
	rssi        int
	address     string
	name        string
	expectedErr error
}

func TestCreateDevice(t *testing.T) {
	testTable := []createDeviceTestTable{{
		rssi:        1,
		address:     "5F:DE:1B:99:C2:ED",
		name:        "test bluetooth",
		expectedErr: nil,
	}, {
		address:     "invalid",
		expectedErr: ErrDeviceAddress,
	}}

	for _, row := range testTable {
		device, err := CreateDevice(row.rssi, row.address, row.name)
		if !errors.Is(err, row.expectedErr) {
			t.Errorf("Expected %v got %v", row.expectedErr, err)
		}
		if device != nil && row.rssi != device.GetRssi() {
			t.Errorf("Expected %d got %d", row.rssi, device.GetRssi())
		}
		if device != nil && row.address != device.GetAddress().String() {
			t.Errorf("Expected %s got %s", row.address, device.GetAddress())
		}
		if device != nil && row.name != device.GetName() {
			t.Errorf("Expected %s got %s", row.name, device.GetName())
		}
	}
}
