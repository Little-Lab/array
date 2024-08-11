package ordermq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"k8s.io/apimachinery/pkg/util/json"
	"order/internal/config"
	"os"
	"zg6/2112a-6/month/model/ordermodel"
)

func Producer(order ordermodel.Orderinfo) {
	var c config.Config
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{c.RocketMQ.NameAddr})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	marshal, err := json.Marshal(order)
	if err != nil {
		return
	}
	msg := &primitive.Message{
		Topic: c.RocketMQ.Topic,
		Body:  marshal,
	}
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
