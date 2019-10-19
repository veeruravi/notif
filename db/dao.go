package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Api interface {
	InsertUrl(url string) (string, error)
	GetUrl(uid string) (string, error)
}

type Dao struct {
	name   string
	db *gorm.DB
}

func NewDao(db *gorm.DB) Api {
	return &Dao{
		name: "UrlDao",
		db:db,
	}
}

func (d *Dao) InsertUrl(url string) (string, error) {
	uid := "123"
	dbUrl := Url{
		Uid: uid,
		Url: url,
	}
	result := d.db.Create(&dbUrl)
	if result.Error != nil {
		return "", errors.New("failed to insert")
	}
	return uid, nil
}

func (d *Dao) GetUrl(uid string) (string, error) {
	var url Url
	result := d.db.Where("uid = ?", uid).Find(&url)
	if result.Error != nil {
		return "", result.Error
	}
	return url.Url, nil
}
