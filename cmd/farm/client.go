package farm

import (
	"fmt"

	"github.com/Pizhlo/whatToFarm-test/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:     "rate",
	Aliases: []string{"client"},
	Short:   "start the client and make get request to server",
	Run: func(cmd *cobra.Command, args []string) {
		pair, _ := cmd.Flags().GetString("pair")

		fmt.Println(pair)

		client.MakeRequest(pair)

	},
}

func init() {
	clientCmd.PersistentFlags().String("pair", "", "receives the necessary pairs of cryptocurrency names (e.g. ETH-USDT)")
	RootCmd.AddCommand(clientCmd)
}
