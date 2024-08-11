package ordermq

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"k8s.io/apimachinery/pkg/util/json"
	"zg6/2112a-6/jobs/model/mysql"
)

const (
	GROUPNAME = "order-consumer-group"
	NAMEADDR  = "127.0.0.1:9876"
	TOPIC     = "order_topic"
)

func Consumer() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(GROUPNAME),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{NAMEADDR})),
	)
	err := c.Subscribe(TOPIC, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		var orderinfo mysql.OrderInfo
		for _, msg := range msgs {
			fmt.Printf("subscribe callback: %v \n", string(msg.Body))
			json.Unmarshal(msg.Body, &orderinfo)
			//归还库存
			err := mysql.UpdateStock(orderinfo.GoodsId, orderinfo.Count)
			if err != nil {
				return 0, errors.New("修改失败:" + err.Error())
			}
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	_ = c.Start()
	select {}
}
