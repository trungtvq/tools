package config_helper

import "github.com/trungtvq/go-utils/config"

func CreateDBConf(host string, port int64, user string, password string, db string) *config.DBInfo{
	return &config.DBInfo{
		Host:     host,
		Port:     port,
		Username: user,
		Password:password,
		DBName:   db,
	}
}
