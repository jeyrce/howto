package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// 发送字符给指定unix域
// 也可以使用: echo "xxx" | nc -U /tmp/test.sock
func TestWriteToUnix(t *testing.T) {
	dial, err := net.Dial("unix", sock)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dial.Close()
	_, err = dial.Write([]byte("测试消息--xxx"))
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Second)
}
