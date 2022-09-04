package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// build connection
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("conn udp server error:%v\n", err)
	}
	defer conn.Close()

	// send data
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		t0 := time.Now()
		_, err = conn.Write([]byte(s)) // send data
		if err != nil {
			fmt.Println("send data failed，err:", err)
			return
		}
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		fmt.Printf("received from server:%v %s\n", string(buf[:n]), time.Since(t0))
	}

}
