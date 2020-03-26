package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; 1 <= 1024; i++ {
		go func(j int) {
			addr := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				//port is closed or filtered
				fmt.Println(addr, "   :CLOSED")
				return
			}
			conn.Close()
			fmt.Println(addr, "   :OPEN")
		}(i)
	}
}
