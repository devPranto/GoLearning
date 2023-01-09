package main

import (
	"fmt"
	"html/template"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/register", Register)
	router.HandleFunc("/submit", Submit).Methods("POST")
	http.ListenAndServe(":8080", router)

}
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, "")

}
func Submit(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hi")
	fmt.Fprint(w, r.FormValue("name"))
}
func Register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/register.html")
	t.Execute(w, nil)

}
