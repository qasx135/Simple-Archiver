package cmd

import "github.com/spf13/cobra"

var vlcCmd = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	Run: 
}

func pack(_ *cobra.Command, args []string) {
	filePath := args[0]
}