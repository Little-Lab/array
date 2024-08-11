package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"zg6/2112a-6/jobs/global"
)

var err error

func InitMysql() {
	c := global.AppConf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DbName)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("mysql 连接失败:" + err.Error())
		return
	}
	log.Println("mysql 连接成功")
}
