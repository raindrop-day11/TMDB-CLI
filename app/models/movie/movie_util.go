package movie

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"tmdb-cli-tool/pkg/config"
	"tmdb-cli-tool/pkg/console"
)

func GetMovies(url string) (movies Movies) {
	//创建一个请求体
	rep, err := http.NewRequest("GET", url, nil)
	console.ExitIf(err)

	//设置请求头
	rep.Header.Set("accept", "application/json")
	rep.Header.Set("Authorization", "Bearer "+config.GetString("tmdb.token"))

	//自定义http请求客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//运行
	rps, err := client.Do(rep)
	console.ExitIf(err)

	defer rps.Body.Close()

	//读取数据
	data, err := io.ReadAll(rps.Body)
	console.ExitIf(err)

	//解析到movies
	err = json.Unmarshal(data, &movies)
	console.ExitIf(err)

	return
}
