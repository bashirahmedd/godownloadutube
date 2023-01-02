package utube

import (
	"fmt"

	"github.com/spf13/cobra"

	cutils "github.com/bashirahmedd/godownloadutube/pkg/common"
	video "github.com/bashirahmedd/godownloadutube/pkg/utube"
)

var reverseCmd = &cobra.Command{

	Use:     "video",
	Aliases: []string{"vid"},
	Short:   "Download videos .mp4",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//res := stringer.Reverse(args[0])
		res := "downloading videos"
		fmt.Println(res)

		config := cutils.GetShareConfig()
		video.Download(config)	
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
