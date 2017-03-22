package main

import (
	"fmt"
	"net"
	"syscall"

	"golang.org/x/net/route"
)

const (
	UGSH = syscall.RTF_UP | syscall.RTF_GATEWAY | syscall.RTF_STATIC | syscall.RTF_HOST
	UGSc = syscall.RTF_UP | syscall.RTF_GATEWAY | syscall.RTF_STATIC | syscall.RTF_PRCLONING
)

func main() {
	if rib, err := route.FetchRIB(syscall.AF_UNSPEC, route.RIBTypeRoute, 0); err == nil {
		if msgs, err := route.ParseRIB(route.RIBTypeRoute, rib); err == nil {
			for _, msg := range msgs {
				m := msg.(*route.RouteMessage)
				if m.Flags == UGSH || m.Flags == UGSc {
					var ip net.IP
					switch a := m.Addrs[syscall.AF_UNSPEC].(type) {
					case *route.Inet4Addr:
						ip = net.IPv4(a.IP[0], a.IP[1], a.IP[2], a.IP[3])
					case *route.Inet6Addr:
						ip = make(net.IP, net.IPv6len)
						copy(ip, a.IP[:])
					}
					fmt.Printf("ip = %s\n", ip)
				}
			}
		}
	}
}
