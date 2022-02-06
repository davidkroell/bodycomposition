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
		flags := cmd.Flags()
		weight, _ := flags.GetFloat64("weight")
		if weight == -1 {
			cmd.PrintErr("No weight specified\n")
			return
		}

		fat, _ := flags.GetFloat64("fat")
		hydration, _ := flags.GetFloat64("hydration")
		bone, _ := flags.GetFloat64("bone")
		boneKg, _ := flags.GetFloat64("bone-mass")
		muscle, _ := flags.GetFloat64("muscle")
		muscleKg, _ := flags.GetFloat64("muscle-mass")
		ts, _ := flags.GetInt64("unix-timestamp")
		visceralFat, _ := flags.GetFloat64("visceral-fat")
		metabolicAge, _ := flags.GetFloat64("metabolic-age")
		physiqueRating, _ := flags.GetFloat64("physique-rating")
		calories, _ := flags.GetFloat64("calories")
		bmi, _ := flags.GetFloat64("bmi")

		bc := bodycomposition.NewBodyComposition(weight, fat, hydration, bone, boneKg, muscle, muscleKg, visceralFat, physiqueRating, metabolicAge, calories, bmi, ts)

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

		if err := bodycomposition.Upload(email, password, bc); err != nil {
			cmd.PrintErrf("Error uploading weight to Garmin Connect: %s\n", err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	flags := uploadCmd.Flags()

	flags.StringP("email", "e", "", "Email of the Garmin account")
	flags.StringP("password", "p", "", "Password of the Garmin account")

	flags.Float64P("weight", "w", -1, "Set your weight in kilograms")
	flags.Float64P("fat", "f", 0, "Set your fat in percent")
	flags.Float64("hydration", 0, "Set your hydration in percent")
	flags.Float64P("bone", "b", 0, "Set your bone mass in percent")
	flags.Float64("bone-mass", 0, "Set your bone mass in kilograms")
	flags.Float64P("muscle", "m", 0, "Set your muscle mass in percent")
	flags.Float64("muscle-mass", 0, "Set your muscle mass in kilograms")
	flags.Float64P("calories", "c", 0, "Set your caloric intake")
	flags.Float64("visceral-fat", 0, "Set your visceral fat rating (valid values: 1-60)")
	flags.Float64("metabolic-age", 0, "Set your metabolic age")
	flags.Float64("physique-rating", 0, "Set your physique rating (valid values: 1-9)")
	flags.Float64("bmi", 0, "Set your BMI - body mass index")
	flags.Int64P("unix-timestamp", "t", -1, "Set the timestamp of the measurement")
}
