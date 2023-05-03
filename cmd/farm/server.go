package farm

import (
	"github.com/spf13/cobra"

	"github.com/Pizhlo/whatToFarm-test/internal/farm"
)

// Command to start server 
var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"server"},
	Short:   "start the server",
	Run: func(cmd *cobra.Command, args []string) {
		port := 3001
		farm.Server(port)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
