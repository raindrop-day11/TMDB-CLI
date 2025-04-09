package main

import (
	"fmt"
	"tmdb-cli-tool/app/cmd"

	configbts "tmdb-cli-tool/config"
	"tmdb-cli-tool/pkg/config"
	"tmdb-cli-tool/pkg/console"

	"github.com/spf13/cobra"
)

func init() {
	configbts.InitLize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "tedb-app",
		Short: "just receive a parameter: type, you can get the type movies",

		//所有子命令都会执行此代码
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			//初始化配置项
			config.InitConfig()
		},
	}

	rootCmd.AddCommand(
		cmd.Serve,
	)

	cmd.RegisterDefault(rootCmd, cmd.Serve)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(console.Red("Failed to run the app with %s", err.Error()))
	}
}
