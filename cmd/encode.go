package cmd

import (
	"encoding/base64"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encodes the given input into base64",
	Long: `You can pass anything multiline text.
	JSON or any thing`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Getting base64 encoded string... \n\n")
		str := encodeToBase64(args[0])
		log.Printf("%s\n", str)
		if err := clipboard.WriteAll(str); err == nil {
			log.Success("Copied to clipboard.")
		}

	},
}

func encodeToBase64(inp string) string {
	return base64.URLEncoding.EncodeToString([]byte(inp))
}
