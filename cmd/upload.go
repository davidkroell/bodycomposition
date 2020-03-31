package cmd

import (
	"github.com/davidkroell/bodycomposition"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:     "upload",
	Short:   "Upload your body composition values to Garmin Connect",
	Aliases: []string{"u", "add"},
	Run: func(cmd *cobra.Command, args []string) {
		weight, _ := cmd.Flags().GetFloat64("weight")
		if weight == -1 {
			cmd.PrintErr("No weight specified\n")
			return
		}

		fat, _ := cmd.Flags().GetFloat64("fat")
		hydration, _ := cmd.Flags().GetFloat64("hydration")
		bone, _ := cmd.Flags().GetFloat64("bone")
		muscle, _ := cmd.Flags().GetFloat64("muscle")

		bc := bodycomposition.NewBodyComposition(weight, fat, hydration, bone, muscle)

		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			cmd.Print("Password for ", email, ": ")
			bytePasswd, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				cmd.PrintErr("Password input failed\n")
			}
			password = string(bytePasswd)
		}

		bc.UploadWeight(email, password)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringP("email", "e", "", "Email of the Garmin account")
	uploadCmd.Flags().StringP("password", "p", "", "Password of the Garmin account")

	uploadCmd.Flags().Float64P("weight", "w", -1, "Set your weight in kilograms")
	uploadCmd.Flags().Float64P("fat", "f", 0, "Set your fat in percent")
	uploadCmd.Flags().Float64("hydration", 0, "Set your hydration in percent")
	uploadCmd.Flags().Float64P("bone", "b", 0, "Set your bone mass in percent")
	uploadCmd.Flags().Float64P("muscle", "m", 0, "Set your muscle mass in percent")
}
