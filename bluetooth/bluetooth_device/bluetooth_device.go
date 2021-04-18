package bluetooth_device

type BluetoothDevice interface {
	GetRssi() int
	GetAddress() BluetoothDeviceAddress
	GetName() string
}
