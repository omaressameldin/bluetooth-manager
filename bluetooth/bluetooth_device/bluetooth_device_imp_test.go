package bluetooth_device

import "testing"

type createDeviceTestTable struct {
	rssi    int
	address string
	name    string
}

func TestCreateDevice(t *testing.T) {
	testTable := []createDeviceTestTable{{
		rssi:    1,
		address: "5f:de:1b:99:c2:ed",
		name:    "test bluetooth",
	}}

	for _, row := range testTable {
		device := CreateDevice(row.rssi, row.address, row.name)
		if row.rssi != device.GetRssi() {
			t.Errorf("Expected %d got %d", row.rssi, device.GetRssi())
		}
		if row.address != device.GetAddress() {
			t.Errorf("Expected %s got %s", row.address, device.GetAddress())
		}
		if row.name != device.GetName() {
			t.Errorf("Expected %s got %s", row.name, device.GetName())
		}
	}
}
