package config

import (
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/rest"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var Set = wire.NewSet(
	MysqlConfig,
	HttpServerConfig,
)

type DBEnv struct {
	User string
	Pass string
	Addr string
	Name string
}

type ServerEnv struct {
	Port int
}

func MysqlConfig() *mysql.Config {
	cfg := mysql.NewConfig()

	var e DBEnv
	err := envconfig.Process("db", &e)
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg.User = e.User
	cfg.Passwd = e.Pass
	cfg.Net = "tcp"
	cfg.Addr = e.Addr
	cfg.DBName = e.Name
	cfg.ParseTime = true
	return cfg
}

func HttpServerConfig() *rest.ServerConfig {
	var e ServerEnv
	err := envconfig.Process("", &e)
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := &rest.ServerConfig{Port: e.Port}
	return cfg
}
