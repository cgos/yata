package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Todo.",
	Long:  "This commands adds a new Todo to the list stored on the local file system.",
}