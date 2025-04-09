package config

import "tmdb-cli-tool/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"port": config.Env("APP_PORT", 8000),
		}
	})
}
