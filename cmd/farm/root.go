package farm

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:  "Farm",
    Short: "Farm - a simple console application for displaying cryptocurrency exchange rates",
    Long: `Use Farm to see rates of cryptocurrency`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprint(os.Stderr, "Sorry, there was an error while executing your CLI: ", err)
        os.Exit(1)
    }
}