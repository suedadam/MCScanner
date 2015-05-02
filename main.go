package main

import (
        "fmt"
        "log"
        "net"
        "os"
        "encoding/json"
        "github.com/anvie/port-scanner"
        "github.com/geNAZt/minecraft-status/data"
        "github.com/geNAZt/minecraft-status/protocol"
)

func main() {
        if len(os.Args) < 2 {
                fmt.Printf("Usage: %s <CIDR>\n", os.Args[0])
                os.Exit(2)
        }
        ip, ipnet, err := net.ParseCIDR(os.Args[1])
        if err != nil {
                log.Fatal(err)
        }
        for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
                ps := portscanner.NewPortScanner(ip.String())
                MCCheck1 := ps.GetOpenedPort(25565, 25565)

                if (len(MCCheck1) != 0) {

                        conn, err := protocol.NewNetClient(ip.String())
                        if err != nil {
                                return
                        }

                        conn.SendHandshake()
                        conn.State = protocol.Status

                        conn.SendStatusRequest()
                        statusPacket, errPacket := conn.ReadPacket()
                        if errPacket != nil {
                                return
                        }
                        conn.SendClientStatusPing()
                        _, errPingPacket := conn.ReadPacket()
                        if errPingPacket != nil {
                                return 
                        }
                        status := &data.Status{}
                        errJson := json.Unmarshal([]byte(statusPacket.(protocol.StatusResponse).Data), status)
                        if errJson != nil {
                                return
                        }

                        fmt.Printf("%s:%s\n", ip.String(), status.Description)
                }
        }       
}

func inc(ip net.IP) {
        for j := len(ip)-1; j>=0; j-- {
                ip[j]++
                if ip[j] > 0 {
                        break
                }
        }
}