package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8040")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) //新建goroutines处理连接
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	//var  ch =make(chan struct{})
	for input.Scan() {
		wg.Add(1)
		go func(c net.Conn, shout string, delay time.Duration) {
			defer wg.Done()
			fmt.Fprintln(c, "\t", strings.ToUpper(shout))
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", shout)
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", strings.ToLower(shout))
			//ch<-struct{}{}

		}(c, input.Text(), 1*time.Second)
	}
	wg.Wait()
	//cw := c.(*net.TCPConn)
	//cw.CloseWrite()

	c.Close()
}
