package main

import (
	"log"
	"net"
	"strings"
)

func getLocalIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err == nil {
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr).String()
		parts := strings.Split(localAddr, ":")
		ip = parts[0]
		return ip, nil

	}
	log.Println(err)
	return "", err

}
