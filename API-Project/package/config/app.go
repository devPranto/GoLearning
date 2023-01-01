package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "pranto:password@tcp(127.0.0.1:3306)/prayerTime?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("there is a problem ", err)
	}
	db = d

}
func GetDb() *gorm.DB {
	return db
}
