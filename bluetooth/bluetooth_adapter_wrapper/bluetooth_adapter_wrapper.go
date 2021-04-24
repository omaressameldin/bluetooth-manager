package bluetooth

import "bluetooth-manager/bluetooth/bluetooth_device"

type BluetoothAdapterWrapper interface {
	Search() []bluetooth_device.BluetoothDevice
}
