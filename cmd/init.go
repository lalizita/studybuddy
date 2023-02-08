/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/lalizita/studybuddy/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new studybuddy database and table",
	Long:  `Initialise a new studybuddy database and table.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
