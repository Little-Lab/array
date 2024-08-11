package initialize

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"zg6/2112a-6/jobs/global"
)

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://admin:123456@localhost:27017,localhost:27018,localhost:27019")
	global.MongoDB, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 确保连接成功
	err = global.MongoDB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}
