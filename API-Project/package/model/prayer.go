package model

import (
	"api-project/package/config"
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type PrayerTiming struct {
	ID      int `gorm:"primary_key"`
	Date    string
	Fajr    string `json:"Fajr"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&PrayerTiming{})
}
func (p *PrayerTiming) Creat() {
	db.NewRecord(p)
	db.Create(&p)
}
func FindByDate(date int) PrayerTiming {
	var times PrayerTiming
	db.First(&times, date)
	fmt.Println(times)
	return times
}
