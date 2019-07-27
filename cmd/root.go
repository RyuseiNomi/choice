package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "choice",
	Short: "It's a powerful grouping cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		do(),
	)
}
