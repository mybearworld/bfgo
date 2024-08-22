package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/mybearworld/bfgo/pkg/bf"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "bfgo",
	Short:        "Run BF programs. Written in Go!",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Run BF programs. Written in Go!

Run as:
  bfgo ./file.bf
	bfgo -         # From stdin`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		var (
			content []byte
			err     error
		)
		if filename == "-" {
			content, err = io.ReadAll(os.Stdin)
		} else {
			content, err = os.ReadFile(filename)
		}
		if err != nil {
			errorAndExit(err)
		}
		err = bf.Run(content, !noNewLine)
		if err != nil {
			errorAndExit(err)
		}
	},
}

func errorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "bfgo: %v\n", err)
	os.Exit(1)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Usage: bfgo run ./file.bf\n")
		os.Exit(1)
	}
}

var noNewLine bool

func init() {
	rootCmd.Flags().BoolVarP(&noNewLine, "no-new-line", "n", false, "do not print a \\n character after the output")
}
