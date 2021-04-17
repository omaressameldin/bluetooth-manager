package bluetooth

type BluetoothDevice interface {
	getRssi() int
	getAddress() string
	getName() string
}
