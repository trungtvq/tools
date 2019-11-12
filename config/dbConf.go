package config

import (
	"sync"

	"github.com/spf13/viper"
)

type DB struct {
	Db   *DBInfo
	Once sync.Once
}
type DBInfo struct {
	Host     string
	Port     int64
	Username string
	Password string
	DBName   string
}

type EnvDB string

const (
	srcDB EnvDB = "mysql-src"
	desDB EnvDB = "mysql-des"
)

func initDB() {
	GetMySQLSrcClient()
	GetMySQLDesClient()
}

var (
	mySQLSrcInfo DB
	mySQLDesInfo DB
)

// GetMySQLSrcClient to read host, port, username, password, dbname
func GetMySQLSrcClient() *DBInfo {
	mySQLSrcInfo.Once.Do(func() {
		mySQLSrcInfo.Db = readDBConf(srcDB)

	})
	return mySQLSrcInfo.Db
}

// GetMySQLDesClient to read host, port, username, password, dbname
func GetMySQLDesClient() *DBInfo {
	mySQLDesInfo.Once.Do(func() {
		mySQLDesInfo.Db = readDBConf(desDB)
	})
	return mySQLDesInfo.Db
}

func readDBConf(db EnvDB) *DBInfo {
	return &DBInfo{
		Host:     viper.GetString(string(db) + ".host"),
		Port:     viper.GetInt64(string(db) + ".port"),
		Username: viper.GetString(string(db) + ".username"),
		Password: viper.GetString(string(db) + ".password"),
		DBName:   viper.GetString(string(db) + ".dbname"),
	}
}

