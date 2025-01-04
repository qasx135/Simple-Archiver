package cmd

import (
	"archiver/lib/vlc"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
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
	packCmd.AddCommand(vlcCmd)
}
