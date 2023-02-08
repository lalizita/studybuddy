package cmd

import (
	"github.com/lalizita/studybuddy/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See a list of all notes you've added",
	Long:  `See a list of all notes you've added.`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
