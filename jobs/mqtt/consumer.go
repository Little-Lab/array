package mqt

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
	"zg6/2112a-6/jobs/global"
)

const (
	SERVER   = "tcp://127.0.0.1:1883"
	ID       = "mqttx"
	USERNAME = "123456"
	PASSWORD = "123456"
	TOPIC    = "test1"
)

// 消息处理函数
func publishHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf(">>>>>>>>>>>>>>>>>>>MESSAGE : %s\n", msg.Payload())
	fmt.Printf("TOPIC : %s\n", msg.Topic())
	//数据过滤
	var data map[string]interface{}
	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		panic(err)
	}

	//TODO:添加
	//global.MongoDB.Database("wk_health_data").Collection("health_data").InsertOne(context.Background(), data)
	//TODO:删除
	//global.MongoDB.Database("wk_health_data").Collection("health_data").DeleteOne(context.Background(), data)
	//TODO:查询
	//value := data["value"]
	//result := global.MongoDB.Database("wk_health_data").Collection("health_data").FindOne(context.Background(), bson.M{"value": value})
	//var doc bson.M
	//result.Decode(&doc)
	//log.Printf("查询到的信息: %+v", doc)
	//TODO:修改
	global.MongoDB.Database("wk_health_data").Collection("health_data").UpdateOne(context.Background(),
		bson.M{"value": 19},
		bson.M{"$set": bson.M{"value": 23}})

}

func Consumer() {
	// 创建 MQTT 客户端选项
	opts := mqtt.NewClientOptions().AddBroker(SERVER).SetClientID(ID).SetUsername(USERNAME).SetPassword(PASSWORD)
	opts.SetKeepAlive(60 * time.Second)           // 设置心跳间隔
	opts.SetDefaultPublishHandler(publishHandler) // 设置默认的消息处理函数
	opts.SetPingTimeout(1 * time.Second)          // 设置 ping 超时时间

	// 创建 MQTT 客户端
	c := mqtt.NewClient(opts)

	// 连接到 MQTT 代理
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error()) // 处理连接错误
	}

	// 订阅主题
	if token := c.Subscribe(TOPIC, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error()) // 处理订阅错误
		os.Exit(1)
	}
	defer func() {
		// 取消订阅主题
		if token := c.Unsubscribe(TOPIC); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error()) // 处理取消订阅错误
			os.Exit(1)
		}

		// 断开与代理的连接
		c.Disconnect(250)
	}()

	select {}
}
