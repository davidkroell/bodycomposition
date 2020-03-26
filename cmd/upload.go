package cmd

import (
	"github.com/davidkroell/bodycomposition"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload your body values to Garmin connect cloud",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		weight, _ := cmd.Flags().GetFloat64("weight")
		fat, _ := cmd.Flags().GetFloat64("fat")
		hydration, _ := cmd.Flags().GetFloat64("hydration")
		bone, _ := cmd.Flags().GetFloat64("bone")
		muscle, _ := cmd.Flags().GetFloat64("muscle")

		bc := bodycomposition.NewBodyComposition(weight, fat, hydration, bone, muscle)

		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		bc.UploadWeight(email, password)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringP("email", "e", "", "Email of the Garmin account")
	uploadCmd.Flags().StringP("password", "p", "", "Password of the Garmin account")

	uploadCmd.Flags().Float64P("weight", "w", 0, "Set your weight in kilograms")
	uploadCmd.Flags().Float64P("fat", "f", 0, "Set your fat in percent")
	uploadCmd.Flags().Float64("hydration", 0, "Set your hydration in percent")
	uploadCmd.Flags().Float64P("bone", "b", 0, "Set your bone mass in percent")
	uploadCmd.Flags().Float64P("muscle", "m", 0, "Set your muscle mass in percent")
}
