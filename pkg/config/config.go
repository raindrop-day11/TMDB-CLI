package config

import (
	"tmdb-cli-tool/pkg/helpers"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

// 全局Viper实例
var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

// 存储配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	//初始化全局viper对象
	viper = viperlib.New()

	//配置文件类型是什么
	viper.SetConfigType("env")
	//配置文件查找路径,.代表从main.go开始找
	viper.AddConfigPath(".")
	//配置文件内配置名称的前缀是什么
	viper.SetEnvPrefix("appenv")

	//读取环境变量
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig() {
	//加载 配置文件（env）中 的配置信息
	loadEnv()
	//将配置信息 移动到viper内部
	loadConfig()
}

func loadEnv() {
	//env文件名称
	envdir := ".env"

	viper.SetConfigName(envdir)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//配置文件需要自动更新，即监控
	viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func Add(modulename string, confunc ConfigFunc) {
	ConfigFuncs[modulename] = confunc
}

func Env(moduleName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(moduleName, defaultValue...)
	}
	return internalGet(moduleName)
}

func internalGet(moduleName string, defaultValue ...interface{}) interface{} {
	//判断一下viper中是否有这个配置信息，并且不为空。返回这个值
	//如果不是，判断一下是否传入默认值
	//如果是，返回默认值，否则返回nil
	if !viper.IsSet(moduleName) || helpers.IsEmpty(viper.Get(moduleName)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(moduleName)
}

func GetString(moduleName string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(moduleName, defaultValue...))
}

func GetInt(moduleName string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(moduleName, defaultValue...))
}

func GetBool(moduleName string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(moduleName, defaultValue...))
}
