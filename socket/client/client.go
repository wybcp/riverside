package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:8082"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error:%v", err)
	}
}
