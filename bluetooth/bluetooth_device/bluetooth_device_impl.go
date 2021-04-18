package bluetooth_device

type bluetoothDeviceImpl struct {
	rssi    int
	address string
	name    string
}

func CreateDevice(rssi int, address, name string) BluetoothDevice {
	return bluetoothDeviceImpl{
		rssi:    rssi,
		address: address,
		name:    name,
	}
}

func (bd bluetoothDeviceImpl) GetAddress() string {
	return bd.address
}

func (bd bluetoothDeviceImpl) GetName() string {
	return bd.name
}

func (bd bluetoothDeviceImpl) GetRssi() int {
	return bd.rssi
}
