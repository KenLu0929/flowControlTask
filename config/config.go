package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	DB struct {
		Name     string `env:"DBName" default:"flow_control_task"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
	RDB struct {
		Host     string `env:"RDBHost" default:"localhost"`
		Port     string `env:"RDBPort" default:"6379"`
		Password string `env:"RDBPassword"`
		Db       int    `env:"RDBDb" default:"0"`
	}
	JWT struct {
		Secret string `env:"JWTsecret"`
	}
}{}

func init() {
	if err := configor.Load(&Config, "config/config.yml"); err != nil {
		panic(err)
	}
}
