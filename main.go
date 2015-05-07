package main

import (
	"log"
	"net"
	"runtime"
	"os"
	"io"
)

var cpus = runtime.NumCPU()

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %s <CIDR>\n", os.Args[0])
		os.Exit(2)
	}
	runtime.GOMAXPROCS(cpus)
	f, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Can't use output.txt! :(")
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	p := newPool(1000)

	ip, ipnet, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		p.add(ip.String())
	}
	p.end()
	log.Println("Done")
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
