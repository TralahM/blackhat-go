package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("0.0.0.0:%d", p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	for i := 1; 1 < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Println("port  open\n", port)
	}
}
