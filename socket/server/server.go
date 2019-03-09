package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":8082"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error:%v", err)
	}
}
