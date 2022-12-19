package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

// 模拟nsq 生产者
func producer() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	p1, err := nsq.NewProducer("10.20.23.42:4150", nsq.NewConfig())
	p2, err := nsq.NewProducer("10.20.23.42:4152", nsq.NewConfig())
	p3, err := nsq.NewProducer("10.20.23.42:4154", nsq.NewConfig())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer p1.Stop()
	defer p2.Stop()
	defer p3.Stop()
	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			if err = p1.Publish("test", []byte(fmt.Sprintf("%d", count))); err != nil {
				fmt.Println(err.Error())
			}
			if err = p2.Publish("test", []byte(fmt.Sprintf("%d", count))); err != nil {
				fmt.Println(err.Error())
			}
			if err = p3.Publish("test", []byte(fmt.Sprintf("%d", count))); err != nil {
				fmt.Println(err.Error())
			}
		default:
			// pass
		}
	}
}
