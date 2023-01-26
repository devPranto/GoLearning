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
	var imgServer = http.FileServer(http.Dir("./picture/"))
	router.PathPrefix("/picture/").Handler(http.StripPrefix("/picture/", imgServer))
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/register", controllers.Register)
	router.HandleFunc("/submit", controllers.RegisterUser)
	router.HandleFunc("/authenticate", controllers.Authenticate)
	router.HandleFunc("/update", controllers.Update)
	router.HandleFunc("/logout", controllers.Logout)
	router.HandleFunc("/updateInfo", controllers.UpdateInfo).Methods("POST")
	router.HandleFunc("/block/{email}", controllers.Block)
	router.HandleFunc("/decode", controllers.DecodeHash)
	router.HandleFunc("/search", controllers.Search)
	router.HandleFunc("/search/{email}", controllers.ShowBlocks)
	//http.Handle("/", controllers.Middleware(router))
	log.Fatal(http.ListenAndServe(":8080", router))

}
