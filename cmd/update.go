/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"mycobra/helper"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "The 'update' subcommand will update a passed in key value pair for an existing set of data to the application configuration file.",
	Long: `The 'update' subcommand updates a key value pair, if the key value pair already exists it is updated, if it does
not exist then the passed in values are added to the application configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		helper.UpdateKeyValuePair(key, value)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&key, "key", "k", "", "The key to update in the configuration file.")
	updateCmd.Flags().StringVarP(&value, "value", "v", "", "The value to update in the configuration file.")
	updateCmd.MarkFlagRequired("key")
	updateCmd.MarkFlagRequired("value")
}
