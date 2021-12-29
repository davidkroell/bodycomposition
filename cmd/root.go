package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v2.2.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "bodycomposition",
	Long: `Bodycomposition is a program to manage your body measurements and composition stored in
Garmin Connect Cloud (https://connect.garmin.com) from your beloved commandline.

For now, you can just add body composition values. Any other thing should be done in Garmin Connect.
Version ` + version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
