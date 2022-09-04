package main

import (
	"fmt"
	"net"
)

func main() {
	// listen to a port
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("conn failed error:%v\n", err)
		return
	}
	defer conn.Close()

	for {
		// receive data
		var data [1024]byte
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read data error:%v\n", err)
			return
		}
		fmt.Printf("addr:%v\t count:%v\t data:%v\n", addr, n, string(data[:n]))
		// send data
		_, err = conn.WriteToUDP(append(data[:n], " received"...), addr)
		if err != nil {
			fmt.Printf("send data error:%v\n", err)
			return
		}
	}
}
