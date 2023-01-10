package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"login-project/models"
	"os"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/register", Register)
	router.HandleFunc("/submit", Submit).Methods("POST")
	router.HandleFunc("/check", check).Methods("POST")
	http.ListenAndServe(":8080", router)

}

func login(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("views/login.html")
	t.Execute(writer, nil)
}

func check(writer http.ResponseWriter, request *http.Request) {
	Email := strings.TrimSpace(request.FormValue("email"))
	password := strings.TrimSpace(request.FormValue("password"))
	result := models.Find(Email)
	fmt.Println(result)
	err := bcrypt.CompareHashAndPassword(result.Password, []byte(password))
	if err == nil {
		fmt.Fprint(writer, "SUCCESS")
	} else {
		fmt.Fprint(writer, "Failed")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, "")

}
func Submit(w http.ResponseWriter, r *http.Request) {
	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)
	userData := models.User{
		FirstName: strings.TrimSpace(r.FormValue("Fname")),
		LastName:  strings.TrimSpace(r.FormValue("Lname")),
		Email:     strings.TrimSpace(r.FormValue("email")),
		Gender:    r.FormValue("gender"),
		Password:  password,
	}
	file, header, err := r.FormFile("image")
	fmt.Printf("file : %v \n header : %v \n error: %v \n", file, header, err)
	//fmt.Printf("User Picture: %v \n", r.FormValue("image"))
	fmt.Println(userData)
	fmt.Fprint(w, userData)
	userData.Insert()
}
func Register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/register.html")
	t.Execute(w, nil)

}

func database() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("mongodb+srv://prantodev:password18@cluster0.yyh6vb2.mongodb.net/?retryWrites=true&w=majority")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
