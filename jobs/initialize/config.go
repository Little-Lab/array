package initialize

import (
	"github.com/spf13/viper"
	"log"
	"zg6/2112a-6/jobs/global"
)

func InitConfig() {
	path := "../jobs/initialize/config.yaml"
	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		log.Println("动态读取失败")
		return
	}
	viper.Unmarshal(&global.NacosConf)
}
