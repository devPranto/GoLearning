package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":9020", router))

}
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, "")
}