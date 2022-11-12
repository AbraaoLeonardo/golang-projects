package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// ADD YOUR CREDENTIAL {USER, PASSWORD, DATABASE}
	dns := "XXXX:XXXX@tcp(localhost:3306)/XXXX?charset=utf8mb4&parseTime=True&loc=Local"
	Connected, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	db = Connected
}

func GetDB() *gorm.DB {
	return db
}
