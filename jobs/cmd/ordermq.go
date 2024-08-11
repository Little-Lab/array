/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"zg6/2112a-6/jobs/rocketmq/ordermq"
)

// ordermqCmd represents the ordermq command
var ordermqCmd = &cobra.Command{
	Use:   "ordermq",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: UpdateStock,
}

func init() {
	rootCmd.AddCommand(ordermqCmd)
}
func UpdateStock(*cobra.Command, []string) {
	ordermq.Consumer()
}
