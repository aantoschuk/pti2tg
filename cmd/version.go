package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func VersionHandler() {
	data, err := os.ReadFile("VERSION")
	if err != nil {
		fmt.Println("Could not read version", err)
		return
	}

	version := strings.TrimSpace(string(data))
	fmt.Println(version)
}

// help command
var VersionCmd = &cobra.Command{
	Use:                   "v",
	Short:                 "Print the version of the pti2tg",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		VersionHandler()
	},
}
