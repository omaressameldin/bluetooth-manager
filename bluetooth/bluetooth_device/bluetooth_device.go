package bluetooth_device

type BluetoothDevice interface {
	GetRssi() int
	GetAddress() string
	GetName() string
}
