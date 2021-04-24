package bluetooth

import (
	"bluetooth-manager/bluetooth/bluetooth_adapter_wrapper/bluetooth_adapter"
	"bluetooth-manager/bluetooth/bluetooth_device"
	"bluetooth-manager/utils"
	"time"

	bluetooth_connector "tinygo.org/x/bluetooth"
)

const scanDuration = 5 * time.Second

type bluetoothAdapterWrapperImpl struct {
	adapter        bluetooth_adapter.BluetoothAdapter
	scannedDevices []bluetooth_device.BluetoothDevice
}

func CreateBluetoothAdapter(adapter bluetooth_adapter.BluetoothAdapter) BluetoothAdapterWrapper {
	return bluetoothAdapterWrapperImpl{
		adapter:        adapter,
		scannedDevices: []bluetooth_device.BluetoothDevice{},
	}
}

func (ba bluetoothAdapterWrapperImpl) Search() []bluetooth_device.BluetoothDevice {
	ba.enableAdapter()
	go ba.disableAdapterAfterTimeout()

	err := ba.adapter.Scan(ba.createDeviceFoundFn())
	utils.PanicOnError("start scan", err)

	return ba.scannedDevices
}

func (ba *bluetoothAdapterWrapperImpl) enableAdapter() {
	utils.PanicOnError("enable BLE stack", ba.adapter.Enable())
}

func (ba *bluetoothAdapterWrapperImpl) disableAdapterAfterTimeout() {
	time.Sleep(scanDuration)
	utils.PanicOnError("stop scan", ba.adapter.StopScan())
}

func (ba *bluetoothAdapterWrapperImpl) createDeviceFoundFn() func(*bluetooth_connector.Adapter, bluetooth_connector.ScanResult) {
	deviceFoundFn := func(adapter *bluetooth_connector.Adapter, result bluetooth_connector.ScanResult) {
		newDevice, err := bluetooth_device.CreateDevice(
			int(result.RSSI),
			result.Address.String(),
			getDeviceNameFromResult(result),
		)
		if err == nil {
			ba.scannedDevices = append(ba.scannedDevices, newDevice)
		}

	}

	return deviceFoundFn
}

func getDeviceNameFromResult(result bluetooth_connector.ScanResult) string {
	name := ""
	if result.AdvertisementPayload != nil {
		name = result.LocalName()
	}

	return name
}
