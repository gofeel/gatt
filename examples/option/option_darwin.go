package option

import "github.com/gofeel/gatt"

var DefaultClientOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.CentralManager),
}

var DefaultServerOptions = []gatt.Option{
	gatt.MacDeviceRole(gatt.PeripheralManager),
}
