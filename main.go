package main

import (
	bluetooth "bluetooth-manager/bluetooth/bluetooth_adapter_wrapper"

	bluetooth_connector "tinygo.org/x/bluetooth"

	"fmt"
)

func main() {
	adapter := bluetooth.CreateBluetoothAdapter(bluetooth_connector.DefaultAdapter)
	devices := adapter.Search()
	fmt.Println(len(devices))
	for _, device := range devices {
		fmt.Printf("foudn device %s, rssi: %d, address: %s\n", device.GetName(), device.GetRssi(), device.GetAddress())
	}

}
