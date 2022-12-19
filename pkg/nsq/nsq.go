package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	go producer()
	c, err := nsq.NewConsumer("test", "1", nsq.NewConfig())
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	c.AddHandler(newHandler(os.Stdout))
	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		// 此处不做处理直接丢弃
		return nil
	}))
	// if err := c.ConnectToNSQDs([]string{"10.20.23.42:4150", "10.20.23.42:4152", "10.20.23.42:4154"}); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	if err := c.ConnectToNSQLookupd("10.20.23.42:4160"); err != nil {
		fmt.Println(err.Error())
		return
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	time.Sleep(time.Second)
	c.Stop()
}
