package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")

	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r) // getting slice of bytes
	if err != nil {
		handleErr(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	//----------simple mod---------\\

	// // path = /path/to/file/myFile.txt
	// fileName := filepath.Base(path) // fileName = myFile.txt
	// ext := filepath.Ext(fileName) // ext = .txt
	// baseName := strings.TrimSuffix(fileName, ext) // baseName = myFile

	// return baseName + "." + packedExtension

	//----------work mod---------\\

	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension

}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")
	packCmd.MarkFlagRequired("method")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
