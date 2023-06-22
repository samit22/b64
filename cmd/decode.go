package cmd

import (
	"encoding/base64"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes the given base64 input into string",
	Long:  `If the string is not base64 encoded it errors out.`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Decoding string... \n\n")
		str, err := decodeFromBase64(args[0])
		if err != nil {
			log.Errorf("failed to decode input, err: %v\n", err)
			return
		}
		log.Printf("%s\n", str)
		if err := clipboard.WriteAll(str); err == nil {
			log.Success("Copied to clipboard.\n")
		}

	},
}

func decodeFromBase64(inp string) (string, error) {
	op, err := base64.URLEncoding.DecodeString(inp)
	if err != nil {
		return "", err
	}
	return string(op), nil
}
