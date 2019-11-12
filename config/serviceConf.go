package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	serviceConfig ServiceConfig
	once          sync.Once
	cfgFile       string
)

type ServiceConfig struct {
	ServiceName string
	Description string
}

func GetServiceConfig() *ServiceConfig {
	once.Do(func() {
		serviceConfig.ServiceName = "myService"
		serviceConfig.Description = "Description of my service"
	})
	return &serviceConfig
}
func setDefaultConfig() {
	//set default for log config
	viper.SetDefault("log.path", "/data/logs/go-utils")
	viper.SetDefault("log.name", "debug")
	viper.SetDefault("log.level", "debug")
}
func initServiceConfig() {
	setDefaultConfig()
	//GetServiceConfig()

}
