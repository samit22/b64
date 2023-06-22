/*
Copyright © 2023 Samit
*/
package cmd

import (
	"os"

	"github.com/samit22/b64/logger"
	"github.com/spf13/cobra"
)

var log *logger.Logger

var rootCmd = &cobra.Command{
	Use:   "b64",
	Short: "Encode decode base64 texts",
	Long: `When your secrets are shared in base64 format or 
	you need to convert some string to base64 it help`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Welcome to base64 encoder/decoder!\n")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log = &logger.Logger{}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(encodeCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(decodeCmd)
}
