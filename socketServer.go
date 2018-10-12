package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	var sendTimes = 0

	var port = flag.String("port", "60000", "Specify open port.")
	flag.Parse()

	var ADDRESS = net.JoinHostPort("127.0.0.1", *port)
	receiveBuf := make([]byte, 10, 16)
	listener, err := net.Listen("tcp", ADDRESS)
	dealErr(err)
	fmt.Printf("Listener network address: %v\n", listener.Addr().String())
	conn, err := listener.Accept()
	dealErr(err)
	fmt.Printf("Remote network address: %v\n", conn.RemoteAddr().String())
	// 从client端发送字节，然后读取client返回的字节，就这样循环10次
	for sendTimes < 10 {
		numBytesWritten, err := conn.Write([]byte("Yes."))
		dealErr(err)
		fmt.Printf("On %v, written %v bytes.\n", time.Now().String(), numBytesWritten)
		time.Sleep(1 * time.Second)
		numBytesRead, err := conn.Read(receiveBuf)
		dealErr(err)
		fmt.Printf("On %v, received %v bytes: %v\n", time.Now().String(), numBytesRead, string(receiveBuf[:numBytesRead]))
		sendTimes++
	}
	numBytesWritten, err := conn.Write([]byte("exit"))
	dealErr(err)
	fmt.Printf("On %v, written %v bytes.\n", time.Now().String(), numBytesWritten)
	conn.Close()
	fmt.Println("Done.")
}

func dealErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
