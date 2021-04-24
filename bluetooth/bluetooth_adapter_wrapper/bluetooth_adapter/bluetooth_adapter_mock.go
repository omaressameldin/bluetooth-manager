package bluetooth_adapter

import bluetooth_connector "tinygo.org/x/bluetooth"

type bluetoothAdapterMock struct {
	ScanResults        []bluetooth_connector.ScanResult
	ScanError          error
	EnableAdapterError error
	StopScanError      error
}

func CreateBluetoothAdapterMock(
	scanResults []bluetooth_connector.ScanResult,
	scanError, enableError, stopScanError error,
) BluetoothAdapter {
	results := scanResults
	if results == nil {
		results = []bluetooth_connector.ScanResult{}
	}

	return bluetoothAdapterMock{
		ScanResults:        results,
		ScanError:          scanError,
		EnableAdapterError: enableError,
		StopScanError:      stopScanError,
	}
}

func (ba bluetoothAdapterMock) Scan(callabck func(*bluetooth_connector.Adapter, bluetooth_connector.ScanResult)) error {
	if ba.ScanError != nil {
		return ba.ScanError
	}
	for _, result := range ba.ScanResults {
		callabck(nil, result)
	}

	return nil
}

func (ba bluetoothAdapterMock) Enable() error {
	if ba.EnableAdapterError != nil {
		return ba.EnableAdapterError
	}

	return nil
}

func (ba bluetoothAdapterMock) StopScan() error {
	if ba.StopScanError != nil {
		return ba.StopScanError
	}

	return nil
}

type AdvertismentMock struct {
	Name string
}

func (am AdvertismentMock) LocalName() string {
	return am.Name
}

func (am AdvertismentMock) Bytes() []byte {
	return nil
}

func (am AdvertismentMock) HasServiceUUID(bluetooth_connector.UUID) bool {
	return false
}
