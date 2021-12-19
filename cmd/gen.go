package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate rwf compatible YAML file from provided arguments",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func randSeq() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func genFileName() string {
	return fmt.Sprintf("args-%s.yaml", randSeq())
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(
		&file, "outfile", "o", genFileName(), "the name of the output YAML file",
	)

	genCmd.MarkFlagRequired("file")
}
