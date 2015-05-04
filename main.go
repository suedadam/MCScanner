package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <CIDR>\n", os.Args[0])
		os.Exit(2)
	}
	p := newPool(1000)

	ip, ipnet, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		p.add(ip.String())
	}
	p.end()
	fmt.Println("Done")
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
