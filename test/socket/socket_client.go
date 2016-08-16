package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("tcpAddr: ", tcpAddr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("hello server\n"))
	fmt.Println("send hello to server")
	checkError(err)
	result, err := bufio.NewReader(conn).ReadString('\n')
	checkError(err)
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s\trevieve from server: %s\n", t, result)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
