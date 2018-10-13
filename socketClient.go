package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var port string

func init() {
	flag.StringVar(&port, "port", "60000", "Specify port to connect")
}

func main() {
	flag.Parse()
	var ADDRESS = net.JoinHostPort("127.0.0.1", port)
	receiveBuf := make([]byte, 10, 16)
	conn, err := net.Dial("tcp4", ADDRESS)
	dealErr(err)
	fmt.Printf("Local network address: %v\n", conn.LocalAddr().String())
	fmt.Printf("Remote network address: %v\n", conn.RemoteAddr().String())
	for {
		numBytesRead, err := conn.Read(receiveBuf)
		dealErr(err)
		fmt.Printf("On %v, read %v bytes: %v\n", time.Now().String(), numBytesRead, string(receiveBuf[:numBytesRead]))
		if bytes.Equal(bytes.ToLower(receiveBuf[:numBytesRead]), []byte("exit")) {
			goto dealExit
		}
		time.Sleep(1 * time.Second)
		numBytesWritten, err := conn.Write([]byte("Happy."))
		dealErr(err)
		fmt.Printf("On %v, written %v bytes.\n", time.Now().String(), numBytesWritten)
	}
dealExit:
	conn.Close()
	fmt.Println("Done.")
}

func dealErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
