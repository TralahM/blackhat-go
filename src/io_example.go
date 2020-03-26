package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("IN >>> ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("OUT >>> ")
	return os.Stdout.Write(b)
}

// func Copy(dst io.Writer, src io.Reader) (written int64, error)
func main() {
	//instantiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)
	//create buffer to hold Input/output
	input := make([]byte, 4096)
	//use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Println("Read %d bytes from stdin\n", s)
	//use writer to read output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Println("Wrote %d bytes from stdout\n", s)
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
