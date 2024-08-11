/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"zg6/2112a-6/jobs/cmd"
	"zg6/2112a-6/jobs/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitNacos()
	initialize.InitMysql()
	initialize.InitMongoDB()
	cmd.Execute()
}
