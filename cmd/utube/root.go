package utube

import (
	"fmt"
	// "log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "utubedownloader",
	Short: "utubedownloader - a simple CLI to download video mp4 and playlist details",
	Long: `utubedownloader is a simple CLI (enjoy)
	  
   One can use this application to download videos or playlist detail from the terminal`,

	Run: func(cmd *cobra.Command, args []string) {		
		fmt.Println("Please use appropriate command or enter help.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
