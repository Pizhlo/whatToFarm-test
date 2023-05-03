package farm

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/Pizhlo/whatToFarm-test/pkg/farm"
)

var serverCmd = &cobra.Command{
	Use:     "start server",
	Aliases: []string{"server"},
	Short:   "start the server",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Fprint(os.Stderr, "Unable to convert string to int in Server call:", err)
			os.Exit(1)
		}

		farm.Server(port)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
