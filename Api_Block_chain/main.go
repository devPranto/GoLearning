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

	router.HandleFunc("/login", controllers.Authenticate)
	router.HandleFunc("/register", controllers.RegisterUser)
	router.HandleFunc("/logout", controllers.Logout)
	router.HandleFunc("/update", controllers.Update)
	router.HandleFunc("/updateInfo", controllers.UpdateInfo).Methods("POST")

	router.HandleFunc("/block/{email}", controllers.Block)
	router.HandleFunc("/decode", controllers.DecodeHash)

	router.HandleFunc("/search/{email}", controllers.ShowBlocks)
	log.Fatal(http.ListenAndServe(":8080", router))

}
