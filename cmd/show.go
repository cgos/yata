package cmd

import (
	"fmt"
	"strings"

	"github.com/cgos/yata/internal"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show todos",
	Long:  "This commands show all todos stored on the local file system.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, ""))
		fs := internal.NewFileStore("")
		internal.Add(nil, fs)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
