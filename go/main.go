package main

import (
	"fmt"
	"net"

	"github.com/vpn-kill-switch/sandbox/go/nettest"
)

func main() {
	rifs := nettest.RoutedInterface("ip", net.FlagUp|net.FlagBroadcast)
	if rifs != nil {
		fmt.Printf("Name: %s\n", rifs.Name)
		ip, err := rifs.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Printf("IP: %s\n", ip)
		fmt.Printf("MAC: %s\n", rifs.HardwareAddr.String())
		fmt.Printf("Flags: %s\n", rifs.Flags.String())
	}
}
