package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("Begin listening on 127.0.0.1:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	req, err := bufio.NewReader(conn).ReadString('\n')
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s\trecieve from client: %s\n", t, req)
	checkError(err)
	conn.Write([]byte("welcome, client\n"))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
