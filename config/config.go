package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	DBConfig  DBConfigProp
	RDConfig  RedisConfigProp
	ERRConfig ErrorConfigProp
)

type ErrorConfigProp map[string]string

type DBConfigProp struct {
	Host        string
	Port        string
	User        string
	Pass        string
	Dbname      string
	MaxPool     int
	IdlePool    int
	MaxLifetime time.Duration
}

type RedisConfigProp struct {
	Host  string
	Port  string
	Pass  string
	Index int
}

func initErrConfig() {
	ERRConfig = viper.GetStringMapString("errors")
}

func initDbConfigProp() {
	DBConfig.Host = viper.GetString("db.host")
	DBConfig.Port = viper.GetString("db.port")
	DBConfig.User = viper.GetString("db.user")
	DBConfig.Pass = viper.GetString("db.pass")
	DBConfig.Dbname = viper.GetString("db.dbname")
	DBConfig.MaxPool = viper.GetInt("db.pool.max")
	DBConfig.IdlePool = viper.GetInt("db.pool.idle")
	DBConfig.MaxLifetime = viper.GetDuration("db.pool.lifetime")
}

func initRdConfigProp() {
	RDConfig.Host = viper.GetString("redis.host")
	RDConfig.Port = viper.GetString("redis.port")
	RDConfig.Pass = viper.GetString("redis.pass")
	RDConfig.Index = viper.GetInt("redis.index")
}

func initializeViper() {

	viper.AddConfigPath(".")      // file path
	viper.SetConfigName("config") // file name
	viper.SetConfigType("yaml")   // file extension

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func init() {
	initializeViper()
	initDbConfigProp()
	initRdConfigProp()
	initErrConfig()
}
