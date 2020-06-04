package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Todo.",
	Long:  "This commands adds a new Todo to the list stored on the local file system.",
	Run: func(cmd *cobra.Command, args []string) {
		// yata := main.NewYata()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
