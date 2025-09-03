package config

import (
	"github.com/jinzhu/gorm" // orm
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// the _ is here because that's how it was given in the gorm documentation
)

var (
	db * gorm.DB
	// returns a variable called db which will help other files to interact with the database
)

func Connect() {
	// opens connection with database
	d, err := gorm.Open("mysql", "srishtisahi: <password>@/<tablename>?charset=utf8&parseTime=True&loc=Local")
	// this long login stuff is required by mysql, dont forget to add the charset and everything
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}