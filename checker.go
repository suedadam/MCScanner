package main

import (
	"encoding/json"
	"github.com/geNAZt/minecraft-status/data"
	"github.com/geNAZt/minecraft-status/protocol"
	"net"
	"log"
	"time"
)

func isMinecraft(ip string) bool {
	if !portOpen(ip + ":25565") {
		return false
	}

	conn, err := protocol.NewNetClient(ip)
	if err != nil {
		return false
	}
	defer conn.Close()

	conn.SendHandshake()
	conn.State = protocol.Status

	conn.SendStatusRequest()
	statusPacket, err := conn.ReadPacket()
	if err != nil {
		return false
	}
	status := &data.Status{}
	errJson := json.Unmarshal([]byte(statusPacket.(protocol.StatusResponse).Data), status)
	if errJson != nil {
		return false
	}
	// ScanRes := fmt.Sprintf("%s:%s\n", ip.String(), status.Description)
	log.Printf("%s:%s\n", ip, status.Description)
	return true
}

func portOpen(addr string) bool {
	conn, err := net.DialTimeout("tcp", addr, time.Duration(500*time.Millisecond))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
