package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func readConfigFile(cfgFile *string) {
	if *cfgFile != "" {
		viper.SetConfigFile(*cfgFile)
	} else {
		viper.SetConfigName("config.dev")
	}
	viper.AddConfigPath(".") // Path
	viper.AutomaticEnv()
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config", err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	//watch config
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed", zap.Any("event", e.Name))
	})
}

//InitConfig with file, default is config.dev.toml if it is NOT been set
func InitConfig(cfgFile *string) {
	readConfigFile(cfgFile)
	initServiceConfig()
	//initDBConfig()
}
