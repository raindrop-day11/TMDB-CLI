package config

import "tmdb-cli-tool/pkg/config"

func init() {
	config.Add("tmdb", func() map[string]interface{} {
		return map[string]interface{}{
			"api_key": config.Env("TMDB_APIKEY"),
			"token":   config.Env("TMDB_TOKEN"),
		}
	})
}
