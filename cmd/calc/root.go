package calc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calc",
	Short: "calc - a simple CLI calculator",
	Long: `calc is built in Go using Cobra,
	source code: https://github.com/prodip0609/calc`,

	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Looks like there has been an error while execute your command %s", err)
		os.Exit(1)
	}
}
