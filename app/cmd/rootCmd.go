package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"tmdb-cli-tool/app/models/movie"
	"tmdb-cli-tool/pkg/console"

	"github.com/spf13/cobra"
)

func RunByCmd(cmd *cobra.Command, args []string) {
	var movietype string

	switch Type {
	case "playing":
		movietype = "now_playing"
	case "upcoming":
		movietype = "upcoming"
	case "popular":
		movietype = "popular"
	case "top":
		movietype = "top_rated"
	case "null":
		fmt.Println(console.Red("\n--->  please input the parameter from one of the following: playing, popular, top, upcoming\n"))
		return
	default:
		fmt.Println(console.Red("\n--->  the parameter can only be one of the following: playing, popular, top, upcoming\n"))
		return
	}

	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?language=en-US&page=1", movietype)
	data := movie.GetMovies(url)

	//创建一个用于存储格式化JSON的缓冲区
	var prettyJSON bytes.Buffer

	//格式化JSON（前缀为空字符串，缩进使用4个空格）
	err := json.Indent(&prettyJSON, data, "", "    ")
	console.ExitIf(err)

	//输出到标准输出
	fmt.Println(console.Yellow("\n----------------------------------------------------------------------\n"))

	prettyJSON.WriteTo(os.Stdout)

	fmt.Println(console.Yellow("\n----------------------------------------------------------------------\n"))
}
