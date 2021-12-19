package cmd

import "github.com/spf13/cobra"

var file string

var rootCmd = &cobra.Command{
	Use:     "rwf",
	Version: "1.0",
	Short:   "run tools with lots of flags from a YAML file",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		return
	}
}
