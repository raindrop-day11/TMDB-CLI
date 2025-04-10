package cmd

import (
	"fmt"
	"tmdb-cli-tool/pkg/console"

	"github.com/spf13/cobra"
)

func RunByCmd(cmd *cobra.Command, args []string) {
	switch Type {
	case "playing":
	case "popular":
	case "top":
	case "upcoming":
	case "null":
		fmt.Println(console.Red("\n--->  please input the parameter from one of the following: playing, popular, top, upcoming\n"))
	default:
		fmt.Println(console.Red("\n--->  the parameter can only be one of the following: playing, popular, top, upcoming\n"))
	}
}
