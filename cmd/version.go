package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// read and print version form the file
func VersionHandler() {
	data, err := os.ReadFile("VERSION")
	if err != nil {
		fmt.Println("Could not read version", err)
		return
	}

	version := strings.TrimSpace(string(data))
	fmt.Println(version)
}

var VersionCmd = &cobra.Command{
	Use:                   "v",
	Short:                 "Print the version of the pti2tg",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		VersionHandler()
	},
}
