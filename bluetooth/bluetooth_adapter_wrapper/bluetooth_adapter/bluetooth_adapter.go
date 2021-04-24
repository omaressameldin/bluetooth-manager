package bluetooth_adapter

import bluetooth_connector "tinygo.org/x/bluetooth"

type BluetoothAdapter interface {
	Scan(callabck func(*bluetooth_connector.Adapter, bluetooth_connector.ScanResult)) error
	Enable() error
	StopScan() error
}
