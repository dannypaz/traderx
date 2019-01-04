package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rpcAddress string

var rootCmd = &cobra.Command{
	Use:   "traderx",
	Short: "Traderx is a CLI utility for the sparkswap broker",
	Long:  "A Fast and Flexible CLI utility for the sparkswap broker\nBuilt with love by dannypaz",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute processes the main endpoint of the cobra CLI where we execute the
// rootCmd that triggers all other commands in the cmd folders
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rpcAddress, "rpc-address", "", "rpc address of the specific broker daemon")
}
