package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"errors"
	"fmt"

	"github.com/KenLu0929/flowControlTask/config"
)

var DB *gorm.DB

func init() {
	var err error
	var dsn string
	dbConfig := config.Config.DB

	switch dbConfig.Adapter {
	case "mysql":
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		panic(errors.New("this database adapter not support"))
	}

	if err != nil {
		panic(err)
	}
}
