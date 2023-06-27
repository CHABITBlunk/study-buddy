package cmd

import (
  "github.com/chabitblunk/study-buddy/data"
  "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
  Use: "list",
  Short: "See a list of all notes you've added",
  Long: `See a list of all notes you've added`,
  Run: func (cmd *cobra.Command, args []string) {
    data.DisplayAllNotes()
  },
}

func init() {
  noteCmd.AddCommand(listCmd)
}
