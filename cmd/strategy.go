package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Strategy struct {
	Name             string   `json:"name"`
	Index            []string `json:"index"`
	BuyThresholdPct  float64  `json:"buy_threshold_pct"`
	SellThresholdPct float64  `json:"sell_threshold_pct"`
	BuyIncrement     float64  `json:"buy_increment,omitempty"`
	BuyIncrementPct  float64  `json:"buy_increment_pct,omitempty"`
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
