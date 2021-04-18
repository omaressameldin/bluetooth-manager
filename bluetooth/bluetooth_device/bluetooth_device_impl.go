package bluetooth_device

type bluetoothDeviceImpl struct {
	rssi    int
	address BluetoothDeviceAddress
	name    string
}

func CreateDevice(rssi int, address, name string) (BluetoothDevice, error) {
	deviceAddress, err := CreateBluetoothDeviceAddress(address)
	if err != nil {
		return nil, err
	}

	return &bluetoothDeviceImpl{
		rssi:    rssi,
		address: deviceAddress,
		name:    name,
	}, nil
}

func (bd bluetoothDeviceImpl) GetAddress() BluetoothDeviceAddress {
	return bd.address
}

func (bd bluetoothDeviceImpl) GetName() string {
	return bd.name
}

func (bd bluetoothDeviceImpl) GetRssi() int {
	return bd.rssi
}
