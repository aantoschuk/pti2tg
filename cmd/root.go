package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var imagePath = ""

// starter command
var rootCmd = &cobra.Command{
	Use:   "pti2tg",
	Short: "pti2tg parse text from the image into .txt",
	Long:  "Parse Text from Image To plain text to avoid typing text from image manually",
	Run: func(cmd *cobra.Command, args []string) {
		// enable -v flag
		showVersion, _ := cmd.Flags().GetBool("version")

		if showVersion {
			VersionHandler()
		}
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// register all custom comamnds here
func init() {
	rootCmd.AddCommand(ReadCmd, VersionCmd)
	// register -v flag
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version of the pti2tg")
	ReadCmd.Flags().StringVarP(&imagePath, "path", "p", "", "path to the image file")
	// ReadCmd.MarkFlagRequired("path")
}
