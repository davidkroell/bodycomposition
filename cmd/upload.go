package cmd

import (
	"os"
	"syscall"

	"github.com/davidkroell/bodycomposition"
	"golang.org/x/crypto/ssh/terminal"

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
		ts, _ := cmd.Flags().GetInt64("unix-timestamp")
		visceralFat, _ := cmd.Flags().GetFloat64("visceral-fat")
		metabolicAge, _ := cmd.Flags().GetFloat64("metabolic-age")
		physiqueRating, _ := cmd.Flags().GetFloat64("physique-rating")
		calories, _ := cmd.Flags().GetFloat64("calories")
		maxTries, _ := cmd.Flags().GetInt("max-tries")

		bc := bodycomposition.NewBodyComposition(weight, fat, hydration, bone, muscle, visceralFat, physiqueRating, metabolicAge, calories, ts)

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

		cmd.Println("... uploading weight")

		for i := 0; i < maxTries; i++ {
			if ok := bc.UploadWeight(email, password); ok {
				os.Exit(0)
			}
		}

		cmd.Println("exiting after ", maxTries, " tries")
		os.Exit(1)
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
	uploadCmd.Flags().Float64P("calories", "c", 0, "Set your caloric intake")
	uploadCmd.Flags().Float64("visceral-fat", 0, "Set your visceral fat rating (valid values: 1-60)")
	uploadCmd.Flags().Float64("metabolic-age", 0, "Set your metabolic age")
	uploadCmd.Flags().Float64("physique-rating", 0, "Set your physique rating (valid values: 1-9)")
	uploadCmd.Flags().Int64P("unix-timestamp", "t", -1, "Set the timestamp of the measurement")

	uploadCmd.Flags().Int("max-tries", 1, "Set maximum retry count, if error occur in Garmin Connect api")
}
