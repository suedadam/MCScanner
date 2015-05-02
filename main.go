package main

import (
        "fmt"
        "log"
        "net"
        "os"
        "time"
        "encoding/json"
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
                MCCheck1 := portCheck(ip)
                if (MCCheck1 == 0) {

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
func portCheck(Addr net.IP) int {
        AddrNew := fmt.Sprintf("%s", Addr)
        AddrNew += ":25565"
        _, MCCheck1 := net.DialTimeout("tcp", AddrNew, time.Duration(500 * time.Millisecond))
        if MCCheck1 != nil {
                return 1
        } 
        return 0
}
func inc(ip net.IP) {
        for j := len(ip)-1; j>=0; j-- {
                ip[j]++
                if ip[j] > 0 {
                        break
                }
        }
}