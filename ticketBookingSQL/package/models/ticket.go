package models

import (

	"github.com/jinzhu/gorm"
	"github.com/prantodev/go-ticket/package/config"
)

var (
	db *gorm.DB
)

type Ticket struct {
	Name    string `gorm:"primary_key"`
	Details string 
	TicketNo int `gorm:"auto_increment"`
}
type counter struct{
	Available int 
	Sold int
	Current int
}
//var TicketCount = counter{Available: 50,Sold: 50 , Current:0 }
func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Ticket{})
}
func (t *Ticket) CreateTicket() *Ticket {
	db.NewRecord(t)
	db.Create(&t)
	return t

}

func GetAllGuests() []Ticket {
	var Tickets []Ticket
	db.Find(&Tickets)
	return Tickets
}

