package cmd

import (
	"cmp"
	"fmt"

	"github.com/otiai10/gosseract/v2"
	"github.com/spf13/cobra"
)

// read text from the image, accept path as a string
func ReadHandler(imagePath string) {
	client := gosseract.NewClient()

	defer client.Close()

	err1 := client.SetImage(imagePath)
	text, err2 := client.Text()

	if err := cmp.Or(err1, err2); err != nil {
		fmt.Println("Error reading text from image ", err)
		return
	}

	fmt.Println(text)
}

// command: which is selecting one of the paths, hardcoded or by flag
var ReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read text from image",
	Run: func(cmd *cobra.Command, args []string) {
		hardcodedPath := ""
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println("Error getting the file path ", err)
			return
		}

		// if hardcodedPath provided, override the flag path
		if hardcodedPath != "" {
			path = hardcodedPath
		}

		// if both are empty
		if path == "" {
			fmt.Println("Provide a path using -p flag or hardcode due to tesseract error")
			return
		}
		ReadHandler(path)
	},
}
