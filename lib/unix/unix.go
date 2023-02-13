package main

import (
	"fmt"
	"net"
	"os"
)

// 测试创建unix域进行通信

const (
	unix = "unix"
	sock = "/tmp/test.sock"
)

func main() {
	os.Remove(sock)
	listen, err := net.Listen(unix, sock)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Chown(sock, -1, 100)
	os.Chmod(sock, os.ModeSocket|0660)
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handler(conn)
	}
}

// 连接处理
func handler(conn net.Conn) {
	defer conn.Close()
	var buf = make([]byte, 4096)
	size, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("从 %s 接收到消息: %s", conn.LocalAddr(), string(buf[:size]))
}
