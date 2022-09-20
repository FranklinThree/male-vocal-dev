package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

// GlobalConfig 默认全局配置
var GlobalConfig *Config
var SqlConfig *Config

// Init 使用 ./application.yaml 初始化全局配置
func Init() {
	GlobalConfig = &Config{
		viper.New(),
	}
	GlobalConfig.SetConfigName("application")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("./config")

	err := GlobalConfig.ReadInConfig()
	if err != nil {
		logrus.WithField("config", "GlobalConfig").WithError(err).Panicf("unable to read global config")
	}

	GlobalConfig.WatchConfig() // 自动更新配置
	GlobalConfig.OnConfigChange(func(e fsnotify.Event) {
		err := GlobalConfig.ReadInConfig()
		if err == nil {
			logrus.WithField("config", "GlobalConfig").Info("config updated")
		}
	})

	sqlInit()
}

func sqlInit() {
	SqlConfig = &Config{
		viper.New(),
	}
	SqlConfig.SetConfigName("sql")
	SqlConfig.SetConfigType("yaml")
	SqlConfig.AddConfigPath(".")
	SqlConfig.AddConfigPath("./config")

	err := SqlConfig.ReadInConfig()
	if err != nil {
		logrus.WithField("config", "SqlConfig").WithError(err).Panicf("unable to read sql config")
	}

	SqlConfig.WatchConfig() // 自动更新配置
	SqlConfig.OnConfigChange(func(e fsnotify.Event) {
		err := SqlConfig.ReadInConfig()
		if err == nil {
			logrus.WithField("config", "SqlConfig").Info("config updated")
		}
	})

}
