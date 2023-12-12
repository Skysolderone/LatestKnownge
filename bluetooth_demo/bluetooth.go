package main

import "log"

var adapter = bluetooth.DefaultAdapter()
//参考tinyGo得bluetooth
func main() {
	must("enable BLE interface", adapter.Enable())
	//then
}

func must(action string, err error) {
	if err != nil {
		log.Fatalf("failed to %s,%v", action, err)
	}

}

//scan device
func scanDevice() {
	log.Println("scanning....")
	adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		log.Println("found the device:", device.Address.String(), "RSSI:", device.RSSI)
	})
}
