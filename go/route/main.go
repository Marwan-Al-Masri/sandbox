package main

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/net/route"
)

const UGSH = syscall.RTF_UP | syscall.RTF_STATIC | syscall.RTF_GATEWAY | syscall.RTF_HOST

func main() {
	rib, err := route.FetchRIB(syscall.AF_INET, route.RIBTypeRoute, 0)
	if err != nil {
		log.Fatal(err)
	}
	m, err := route.ParseRIB(route.RIBTypeRoute, rib)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range m {
		x := v.(*route.RouteMessage)
		if x.Flags == UGSH {
			for _, a := range x.Addrs {
				switch t := a.(type) {
				case *route.Inet4Addr:
					fmt.Printf("IP = %v\n", t.IP)
				}
			}
		}
	}
}
