package cmd

import (
	"errors"
	"fmt"
	"os"
	
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type promptContent struct {
	errorMsg 	string
	label 		string
}

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note can be anything you'd like to study and review",
	Long: `A note can be anything you'd like to study and review`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("note called")
	},
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}
	templates := &promptui.PromptTemplates{
		Prompt: 	"{{ . }}",
		Valid: 		"{{ . | green }} ",
		Invalid: 	"{{ . | red }} ",
		Success: 	"{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label: 	pc.label,
		Templates: templates,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewNote() {
	wordPromptContent := promptContent {
		"Please provide a word.",
		"What word would you like to make a note of?",
	}
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide a definition.",
		fmt.Sprintf("What is the definiton of %s?", word),
	}
	definition := promptGetInput(definitionPromptContent)
}

func init() {
	rootCmd.AddCommand(noteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
