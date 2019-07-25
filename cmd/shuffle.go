package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func shuffle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shuffle",
		Short: "shuffle member",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("shuffled")
		},
	}

	return cmd
}
