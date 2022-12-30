package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
//	"github.com/prantodev/go-ticket/package/config"
	"github.com/prantodev/go-ticket/package/routes"
)

func main(){
	r := mux.NewRouter() 
	routes.RegisterTicketsRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":9010",r))

}