package bluetooth

type BluetoothDeviceImpl struct {
	rssi    int
	address string
	name    string
}

func CreateDevice(rssi int, address, name string) BluetoothDevice {
	return BluetoothDeviceImpl{
		rssi:    rssi,
		address: address,
		name:    name,
	}
}

func (bd BluetoothDeviceImpl) getAddress() string {
	return bd.address
}

func (bd BluetoothDeviceImpl) getName() string {
	return bd.name
}

func (bd BluetoothDeviceImpl) getRssi() int {
	return bd.rssi
}
