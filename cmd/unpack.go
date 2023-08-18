package cmd

import (
	"gorchiver/lib/compression"
	"gorchiver/lib/compression/vlc"
	"gorchiver/lib/compression/vlc/table/shannon_fano"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "decompression method: fano")
	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) < 1 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	decompMethod := cmd.Flag("method").Value.String()

	switch decompMethod {
	case "fano":
		decoder = vlc.NewEncoderDecoder(shannon_fano.NewGenerator())
	default:
		cmd.PrintErr("unknown decoding method")
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

	packed := decoder.Decode(data)

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
