package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	urlDao Api
)

func Initialize() {
	var err error;
	db, err := gorm.Open("mysql", viper.GetString("db.url"))
	if err != nil {
		panic("Failed to connect")
	}
	InitializeUrlModel(db)
	urlDao = NewDao(db)
}


func GetUrlDao() Api {
	return urlDao
}