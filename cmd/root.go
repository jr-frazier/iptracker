package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "IPtracker CLI app",
	Long:  `IPtracker CLI app`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Hello World")
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
