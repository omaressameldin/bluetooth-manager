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

func (bd bluetoothDeviceImpl) getAddress() string {
	return bd.address
}

func (bd bluetoothDeviceImpl) getName() string {
	return bd.name
}

func (bd bluetoothDeviceImpl) getRssi() int {
	return bd.rssi
}
