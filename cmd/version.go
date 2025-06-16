package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// help command
var VersionCmd = &cobra.Command{
	Use:                   "v",
	Short:                 "Print the version of the pti2tg",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}
