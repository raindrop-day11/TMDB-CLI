package cmd

import (
	"os"
	"tmdb-cli-tool/pkg/helpers"

	"github.com/spf13/cobra"
)

var Type string

// 注册默认运行
func RegisterDefault(rootCmd *cobra.Command, defaultCmd *cobra.Command) {
	//自rootCmd开始找子命令（包括rootCmd）
	cmd, _, err := rootCmd.Find(os.Args[1:])
	//找参数
	firstArg := helpers.FirstElement(os.Args[1:])
	//判断如果没有子命令，或者有参数且不是-h 那么就运行默认命令
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{defaultCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}

// 注册Type参数
func RegisterGlobalElement(rootCmd *cobra.Command) {
	rootCmd.Flags().StringVarP(&Type, "type", "t", "null", "it represents the movie genre you want to search for")
}
