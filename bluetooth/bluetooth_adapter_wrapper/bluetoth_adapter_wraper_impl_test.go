package bluetooth

import (
	"bluetooth-manager/bluetooth/bluetooth_adapter_wrapper/bluetooth_adapter"
	"bluetooth-manager/bluetooth/bluetooth_device"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	bluetooth_connector "tinygo.org/x/bluetooth"
)

type searchErrorHandlingTestTable struct {
	scanError         error
	enablAdapterError error
	stopScanError     error
}

func TestsearchErrorHandling(t *testing.T) {
	testTable := []searchErrorHandlingTestTable{
		{scanError: fmt.Errorf("test scan error")},
		{enablAdapterError: fmt.Errorf("test enabl error")},
		{stopScanError: fmt.Errorf("test stop scan error")},
		{},
	}
	for _, row := range testTable {
		testAdapter := bluetooth_adapter.CreateBluetoothAdapterMock(
			nil,
			row.scanError,
			row.enablAdapterError,
			row.stopScanError,
		)
		adapterWraper := CreateBluetoothAdapter(testAdapter)
		mightPanicFn := func() { adapterWraper.Search() }
		if row.scanError != nil || row.enablAdapterError != nil || row.stopScanError != nil {
			assert.Panics(
				t,
				mightPanicFn,
			)
		} else {
			assert.NotPanics(t, mightPanicFn)
		}

	}
}

type searchTestTable struct {
	foundDevices []bluetooth_device.BluetoothDevice
}

func TestSearch(t *testing.T) {
	device1, _ := bluetooth_device.CreateDevice(20, "11:22:33:AA:BB:CC", "test1")
	device2, _ := bluetooth_device.CreateDevice(30, "55:55:8D:C7:D4:62", "test")

	testTable := []searchTestTable{
		{
			foundDevices: []bluetooth_device.BluetoothDevice{device1, device2},
		},
	}
	for _, row := range testTable {
		testAdapter := bluetooth_adapter.CreateBluetoothAdapterMock(
			createScanResults(row.foundDevices),
			nil,
			nil,
			nil,
		)

		adapterWrapper := CreateBluetoothAdapter(testAdapter)
		searchResults := adapterWrapper.Search()
		for i, result := range searchResults {
			expectedDevice := row.foundDevices[i]
			if result.GetRssi() != expectedDevice.GetRssi() {
				t.Errorf(
					"rssi of device %d does not match, expected [%d] got [%d]",
					i,
					expectedDevice.GetRssi(),
					result.GetRssi(),
				)
			}
			if result.GetName() != expectedDevice.GetName() {
				t.Errorf(
					"name of device %d does not match, expected [%s] got [%s]",
					i,
					expectedDevice.GetName(),
					result.GetName(),
				)
			}
			if result.GetAddress().String() != expectedDevice.GetAddress().String() {
				t.Errorf(
					"address of device %d does not match, expected [%s] got [%s]",
					i,
					expectedDevice.GetAddress(),
					result.GetAddress(),
				)
			}
		}
	}
}

func createScanResults(devices []bluetooth_device.BluetoothDevice) []bluetooth_connector.ScanResult {
	scanResults := make([]bluetooth_connector.ScanResult, 0, len(devices))
	for _, device := range devices {
		addr, _ := bluetooth_connector.ParseMAC(device.GetAddress().String())
		resultAddress := bluetooth_connector.Address{bluetooth_connector.MACAddress{MAC: addr}}
		result := bluetooth_connector.ScanResult{
			RSSI:    int16(device.GetRssi()),
			Address: resultAddress,
			AdvertisementPayload: bluetooth_adapter.AdvertismentMock{
				device.GetName(),
			},
		}

		scanResults = append(scanResults, result)
	}

	return scanResults
}
