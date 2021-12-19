package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bin3xish477/rwf/utils"
	"github.com/spf13/cobra"
)

func craftCmd(c string, ta *utils.ToolArgs) (cmd []string) {
	cmd = append(cmd, strings.Split(c, " ")...)

	for _, arg := range ta.Long {
		for name, value := range arg {
			switch value.(type) {
			case bool:
				cmd = append(cmd, fmt.Sprintf("--%s", name))
			case string:
				cmd = append(cmd, fmt.Sprintf("--%s", name), value.(string))
			}
		}
	}

	for _, arg := range ta.Short {
		for name, value := range arg {
			switch value.(type) {
			case bool:
				cmd = append(cmd, fmt.Sprintf("-%s", name))
			case string:
				cmd = append(cmd, fmt.Sprintf("-%s", name), value.(string))
			}
		}
	}

	return
}

var runCmd = &cobra.Command{
	Use:   "run [tool] [tool_args]",
	Short: "run a CLI tool with arguments specified in rwf YAML file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var toolArgs utils.ToolArgs
		utils.ParseArgs(file, &toolArgs)

		fmt.Printf("[%s••%s] %sSetting Environment Variables%s:\n",
			utils.Green, utils.End, utils.Bold, utils.End)
		for _, env := range toolArgs.EnvVars {
			for k, v := range env {
				utils.SetEnv(k, v)
			}
		}

		cmdAsList := craftCmd(strings.Join(args, " "), &toolArgs)
		fmt.Printf("[%s••%s] %sCommand%s:\n  + %s\n",
			utils.Red, utils.End, utils.Bold, utils.End,
			strings.Join(cmdAsList, " "),
		)

		fmt.Printf("[%s••%s] %sResults%s:\n",
			utils.Blue, utils.End, utils.Bold, utils.End)

		finish := make(chan bool, 1)
		now := time.Now()

		go func() {
			utils.ExecuteCmd(cmdAsList[0], cmdAsList[1:])
			finish <- true
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, os.Kill)

		go func() {
			<-c
			fmt.Printf("[%s••%s] %sCaught Ctrl+C, Gracefully Exiting...%s\n",
				utils.Red, utils.End, utils.Bold, utils.End)
			return
		}()

		// blocking
		_ = <-finish

		end := time.Since(now)
		fmt.Printf("[%s••%s] %sFinished in %s%s\n",
			utils.Purple, utils.End, utils.Bold, utils.End, end)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(
		&file, "file", "f", "", "the rwf YAML file with arguments",
	)

	runCmd.MarkFlagRequired("file")
}
