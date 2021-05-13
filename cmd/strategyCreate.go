package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var strategyCreateName, strategyCreateIndex string
var strategyCreateBuyThresholdPct, strategyCreateSellThresholdPct,
	strategyCreateBuyIncrement, strategyCreateBuyIncrementPct float64

var strategyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a strategy",
	Long:  `create & save a strategy to run later.  strategies are read & written from specified config file`,
	Run: func(cmd *cobra.Command, args []string) {
		s := Strategy{
			Name:             strategyCreateName,
			IndexName:        strategyCreateIndex,
			BuyThresholdPct:  strategyCreateBuyThresholdPct,
			SellThresholdPct: strategyCreateSellThresholdPct,
			BuyIncrement:     strategyCreateBuyIncrement,
			BuyIncrementPct:  strategyCreateBuyIncrementPct,
		}
		var c Conf
		c.GetConf()
		c.Strategies = append(c.Strategies, s)
		fmt.Println(c)
		c.WriteConf()
	},
}

func init() {
	strategyCmd.AddCommand(strategyCreateCmd)
	strategyCreateCmd.PersistentFlags().StringVarP(&strategyCreateName, "name", "n", "", "name of strategy")
	strategyCreateCmd.PersistentFlags().StringVarP(&strategyCreateIndex, "index", "i", "SP500", "index name for strategy")
	strategyCreateCmd.PersistentFlags().Float64VarP(&strategyCreateBuyThresholdPct, "buy-threshold-pct", "b", 0.05, "buy threshold pct")
	strategyCreateCmd.PersistentFlags().Float64VarP(&strategyCreateSellThresholdPct, "sell-threshold-pct", "s", 0.05, "sell threshold pct")
	strategyCreateCmd.PersistentFlags().Float64Var(&strategyCreateBuyIncrement, "buy-increment", 0, "buy increment")
	strategyCreateCmd.PersistentFlags().Float64Var(&strategyCreateBuyIncrementPct, "buy-increment-pct", 0.25, "buy increment pct")
}
