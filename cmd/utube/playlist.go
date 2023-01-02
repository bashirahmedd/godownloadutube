package utube

import (
	"fmt"

	//"github.com/ThorstenHans/stringer/pkg/stringer"
	"github.com/spf13/cobra"
)

var onlyDigits bool
var inspectCmd = &cobra.Command{

	Use:     "playlist",
	Aliases: []string{"plst"},
	Short:   "Downloads given playlist videos json",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		fmt.Println(onlyDigits)
		//res, kind := stringer.Inspect(i, onlyDigits)
		res, kind := 3, "bashir123"

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(inspectCmd)
}
