package cmd

import (
	"fmt"
	"strings"

	"github.com/cgos/yata/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Todo.",
	Long:  "This commands adds a new Todo to the list stored on the local file system.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, ""))
		fs := internal.NewFileStore("")
		internal.Add(nil, fs)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
