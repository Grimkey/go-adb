package main

import (
	"flag"
	"fmt"
	"log"

	adb "github.com/zhin/go-adb"
)

var (
	port = flag.Int("p", adb.AdbPort, "")

	client *adb.Adb
)

func main() {
	var err error
	client, err = adb.NewWithConfig(adb.ServerConfig{
		Port: *port,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server…")
	client.StartServer()

	serverVersion, err := client.ServerVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println("Server version:", serverVersion)

	devices, err := client.ListDevices()
	if err != nil {
		panic(err)
	}
	fmt.Println("Devices:")
	for _, device := range devices {
		fmt.Printf("\t%+v\n", *device)

		dev := client.Device(adb.DeviceWithSerial(device.Serial))

		vnc := adb.ForwardSpec{
			Protocol:   "tcp",
			PortOrName: "5900",
		}
		err = dev.Reverse(vnc, vnc)
		if err != nil {
			log.Fatal(err)
		}
	}
}
