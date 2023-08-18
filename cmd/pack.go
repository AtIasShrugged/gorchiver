package cmd

import (
	"errors"
	"gorchiver/lib/compression"
	"gorchiver/lib/compression/vlc"
	"gorchiver/lib/compression/vlc/table/shannon_fano"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: fano")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}

const packedExtension = "vlc"
const outputDir = "output/"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) < 1 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	compMethod := cmd.Flag("method").Value.String()

	switch compMethod {
	case "fano":
		encoder = vlc.NewEncoderDecoder(shannon_fano.NewGenerator())
	default:
		cmd.PrintErr("unknown encoding method")
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

	packed := encoder.Encode(string(data))
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0777)
	}
	err = os.WriteFile(filepath.Join(outputDir, packedFileName(filePath)), packed, 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	return baseName + "." + packedExtension
}
