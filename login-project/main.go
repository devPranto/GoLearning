package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	router:= mux.NewRouter()
	router.HandleFunc("/",index).Methods("POST")
	http.ListenAndServe(":8080",router)
	
}
func index(w http.ResponseWriter, r *http.Request)  {
		t,_:= template.ParseFiles("views/index.html")
		t.Execute(w,"")
}