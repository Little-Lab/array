package rmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
)

const (
	GROUPNAME = "consumer-group"
	NAMEADDR  = "127.0.0.1:9876"
	TOPIC     = "sequence_topic"
)

func Consumer() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(GROUPNAME),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{NAMEADDR})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerOrder(true), // 开启顺序消费
	)

	c.Subscribe(TOPIC, consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, m := range msgs {
				log.Printf("Body: %s", string(m.Body))
			}
			return consumer.ConsumeSuccess, nil
		})
	_ = c.Start()
	fmt.Println("开始读取消息 queue")
	select {}
}
