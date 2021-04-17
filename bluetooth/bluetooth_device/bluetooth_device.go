package bluetooth_device

type BluetoothDevice interface {
	getRssi() int
	getAddress() string
	getName() string
}
