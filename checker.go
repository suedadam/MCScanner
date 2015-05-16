package main

import (
	"encoding/json"
	"github.com/suedadam/minecraft-status"
	"net"
	"log"
	"strconv"
	"time"
)

func isMinecraft(ip string, portrec int) bool {
		port := strconv.Itoa(portrec)
		if !portOpen(ip + ":" + port) {
			return false
		}

		conn, err := protocol.NewNetClient(ip + ":" + port)
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
		log.Printf("%s:%s:%s\n", ip, port, status.Description)
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
