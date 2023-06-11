package cmd

import (
	"github.com/chabitblunk/study-buddy/data"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new studybuddy database & table",
	Long: `Initialize a new studybuddy database & table`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
