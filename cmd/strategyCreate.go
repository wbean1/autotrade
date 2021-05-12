package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var strategyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a strategy",
	Long:  `create & save a strategy to run later.  strategies are read & written from specified config file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("strategyCreateCmd called")
	},
}

func init() {
	strategyCmd.AddCommand(strategyCreateCmd)
}
