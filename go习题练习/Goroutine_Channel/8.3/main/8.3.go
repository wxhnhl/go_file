package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8040")
	if err != nil {
		log.Fatal(err)
	}
	//内置make函数创建一个channel,可以发送struct类型的数据
	done := make(chan struct{})
	//go语句调用一个函数字面量，启动goroutine的常用形式
	go func() {
		//从网络连接到标准输出，如果连接没断也会阻塞
		//如果TCP的读连接关闭会报错：use of closed network connection
		_, err := io.Copy(os.Stdout, conn)
		log.Println(err)
		log.Println("done")
		//发送channel给接收goroutine
		done <- struct{}{}
	}()
	//从标准输入到网络连接中，这个地方会阻塞，按Control+D关闭标准输入
	mustCopy(conn, os.Stdin)
	//      conn.Close()
	//类型断言，调用*net.TCPConn的方法CloseWrite()只关闭TCP的写连接
	cw := conn.(*net.TCPConn)
	cw.CloseWrite()
	<-done // 阻塞等待后台 goroutine 完成接收channel
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
