package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"zg6/2112a-6/jobs/config"
)

var (
	AppConf   config.AppConfig
	NacosConf config.NacosConfig
	DB        *gorm.DB
	MongoDB   *mongo.Client
)
