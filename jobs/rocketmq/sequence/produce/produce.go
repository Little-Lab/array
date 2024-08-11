package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

const (
	NAMEADDR = "127.0.0.1:9876"
	TOPIC    = "sequence_topic"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{NAMEADDR})),
		producer.WithRetry(2),
	)
	if err != nil {
		fmt.Printf("NewProducer  error: %s\n", err.Error())
		os.Exit(1)
	}
	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
		os.Exit(1)
	}
	for i := 0; i < 100; i++ {
		msg := &primitive.Message{
			Topic: TOPIC,
			Body:  []byte("======>>>>>>" + strconv.Itoa(i)),
		}
		_, err = p.SendSync(context.Background(), msg) // 不能用单向
		if err != nil {
			log.Printf("send message err: %s", err)
			continue
		}
	}
	log.Println("发布任务")
}
