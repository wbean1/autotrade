package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Strategy struct {
	Name             string  `yaml:"name"`
	IndexName        string  `yaml:"index_name"`
	BuyThresholdPct  float64 `yaml:"buy_threshold_pct"`
	SellThresholdPct float64 `yaml:"sell_threshold_pct"`
	BuyIncrement     float64 `yaml:"buy_increment,omitempty"`
	BuyIncrementPct  float64 `yaml:"buy_increment_pct,omitempty"`
}

var strategyCmd = &cobra.Command{
	Use:   "strategy",
	Short: "do things with strategies",
	Long:  `strategy has subcommands for creating & running strategies`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("strategyCmd called")
	},
}

func init() {
	rootCmd.AddCommand(strategyCmd)
}
