package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {

	// 角色密钥，从 TDMQ 控制台「角色管理」中获取
	authentication := os.Getenv("TDMQ_TOKEN")
	// 接入地址，从 TDMQ 控制台「集群管理」中复制，注意需要包含集群ID
	serviceURL := os.Getenv("TDMQ_URL")

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:              serviceURL,
		Authentication:   pulsar.NewAuthenticationToken(authentication),
		OperationTimeout: 30 * time.Second,
	})
	if err != nil {
		fmt.Println("failed to create client")
		return
	}

	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "persistent://pulsar-mw8mrgxmae37/dbs-cynosdb-gz51/cdp_scheduler_gz51drill_prod",
	})
	if err != nil {
		fmt.Println("failed to create producer", err)
		return
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})

	if err != nil {
		fmt.Println("Failed to publish message", err)
	} else {
		fmt.Println("Published message")
	}

	defer producer.Close()

}
