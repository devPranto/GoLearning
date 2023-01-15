package main

//todo make sure one email cant use multiple times
//todo upload images in local repository and save path in cloud
//fixme handle errors
import (
	"github.com/gorilla/mux"
	"log"
	"login-project/controllers"
	"net/http"
)

//var store = sessions.NewCookieStore(privateKey)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/register", controllers.Register)
	router.HandleFunc("/submit", controllers.Submit)
	router.HandleFunc("/authenticate", controllers.Authenticate)
	router.HandleFunc("/update", controllers.Update)
	router.HandleFunc("/logout", controllers.Logout)
	//http.Handle("/", controllers.Middleware(router))
	log.Fatal(http.ListenAndServe(":8080", router))

}
