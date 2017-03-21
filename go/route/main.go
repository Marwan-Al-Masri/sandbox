package main

import (
	"fmt"
	"log"
	"net"
	"syscall"

	"golang.org/x/net/route"
)

const UGSH = syscall.RTF_UP | syscall.RTF_STATIC | syscall.RTF_GATEWAY | syscall.RTF_HOST

func main() {
	rib, err := route.FetchRIB(syscall.AF_INET, route.RIBTypeRoute, 0)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := route.ParseRIB(route.RIBTypeRoute, rib)
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range msgs {
		m := msg.(*route.RouteMessage)
		if m.Flags == UGSH {
			for _, a := range m.Addrs {
				var ip net.IP
				switch a := a.(type) {
				case *route.Inet4Addr:
					ip = net.IPv4(a.IP[0], a.IP[1], a.IP[2], a.IP[3])
				case *route.Inet6Addr:
					ip = make(net.IP, net.IPv6len)
					copy(ip, a.IP[:])
				}
				fmt.Printf("ip = %+v\n", ip)
				return
			}
		}
	}
}
