package calc

import (
	"fmt"
	"strconv"

	"github.com/prodip0609/calc/pkg/calc"
	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"s"},
	Short:   "Subtract two values",
	Args:    cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		val1, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("command arguments can either be int or float")
		}

		val2, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("command arguments can either be int or float")
		}

		res := calc.Subtract(val1, val2)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
