/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"mycobra/helper"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Thes 'add' subcommand will add a passed in key value pair to the application configuration file.",
	Long: `The 'add' subcommand adds a key value pair to the application configuration file. For example:

'<cmd> add --key theKey --value "the value can be a bunch of things."'.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.AddKeyValuePair(key, value)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&key, "key", "k", "", "The key to add to the configuration file.")
	addCmd.Flags().StringVarP(&value, "value", "v", "", "The value to add to the configuration file.")

	addCmd.MarkFlagRequired("key")
	addCmd.MarkFlagRequired("value")
}
