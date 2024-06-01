/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"mycobra/helper"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "The 'delete' subcommand removes a key value pair from the configuration file.",
	Long:  "The 'delete' subcommand removes a key value pair from the configuration file.",
	Run: func(cmd *cobra.Command, args []string) {
		helper.DeleteKeyValuePair(key)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&key, "key", "k", "", "The key to delete from the configuration file.")
	deleteCmd.MarkFlagRequired("key")
}
