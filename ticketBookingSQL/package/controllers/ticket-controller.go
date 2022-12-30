package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/prantodev/go-ticket/package/models"
	"github.com/prantodev/go-ticket/package/util"
)



func GetGuest(w http.ResponseWriter, r *http.Request) {
	newTicket := models.GetAllGuests()
	t,_ := template.ParseFiles("package/controllers/home.html")
	t.Execute(w,newTicket)

}

func CreateTickets(w http.ResponseWriter, r *http.Request) {
	CreateTicket := &models.Ticket{Name: r.FormValue("name"), Details: r.FormValue("details") }
	util.ParseBody(r, CreateTicket)
	t := CreateTicket.CreateTicket()
	 temp, _ := template.ParseFiles("package/controllers/success.html")
	 temp.Execute(w,t.Name)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at Index")
	t, _ := template.ParseFiles("package/controllers/edit.html")
	count:=50-len(models.GetAllGuests())
	t.Execute(w,count)
}
