package main

import (
	"fmt"
	"io"

	"github.com/nsqio/go-nsq"
)

type handler struct {
	out io.Writer
}

func newHandler(out io.Writer) *handler {
	return &handler{out: out}
}

func (h handler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		return nil
	}
	_, err := fmt.Fprintf(h.out, "%s 收到消息: %s\n", message.NSQDAddress, string(message.Body))
	if err != nil {
		return err
	}
	return nil
}
