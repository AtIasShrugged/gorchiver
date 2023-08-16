package cmd

import (
	"gorchiver/lib/vlc"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) < 1 || args[0] == "" {
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

	packed := vlc.Decode(string(data))

	err = os.WriteFile(filepath.Join(outputDir, unpackedFileName(filePath)), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	return baseName + "." + unpackedExtension
}
