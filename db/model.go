package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Url struct {
	gorm.Model
	Uid string 	`gorm:"column:uid;size:10"`
	Url string  `gorm:"column:url"`
}

func InitializeUrlModel(db *gorm.DB) {
	db.AutoMigrate(&Url{})
	db.Model(&Url{}).AddUniqueIndex("UniqueIndex", "uid")
}