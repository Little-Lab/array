/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	mqt "zg6/2112a-6/jobs/mqtt"
)

// mqttCmd represents the mqtt command
var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: MQTTConsumer,
}

func init() {
	rootCmd.AddCommand(mqttCmd)
}
func MQTTConsumer(*cobra.Command, []string) {
	mqt.Consumer()
}
