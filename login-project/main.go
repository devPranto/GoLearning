package main

//todo make sure one email cant use multiple times
//todo upload images in local repository and save path in cloud
//fixme handle errors
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"login-project/models"
	"net/http"
	"strings"
	"time"
)

var (
	expireyTime   = time.Now().Add(time.Minute * 1)
	privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
)

//var store = sessions.NewCookieStore(privateKey)

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
	// todo uncomment and modify codes in order to check with postman
	//var result Credentials
	//json.NewDecoder(request.Body).Decode(&result)
	//err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte("devhouse"))
	Email := strings.TrimSpace(request.FormValue("email"))
	password := strings.TrimSpace(request.FormValue("password"))
	result := models.Find(Email)
	fmt.Println(result)
	err := bcrypt.CompareHashAndPassword(result.Password, []byte(password))
	if err == nil {
		writer.WriteHeader(http.StatusOK)
		//fmt.Fprint(writer, "SUCCESS")
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		//fmt.Fprint(writer, "Failed")
		return
	}
	token := CreatJWT(result)
	fmt.Println(token)

	http.SetCookie(writer, &http.Cookie{
		Name:    "AHOM",
		Value:   token,
		Expires: expireyTime,
	})
	t, _ := template.ParseFiles("views/welcome.html")
	t.Execute(writer, result)
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

func CreatJWT(user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "nil" //fmt.Sprintf("failed to creat token : %v", err)
	} else {
		return tokenString
	}
}

//func CreatJWT(user *Credentials) string {
//	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
//		"sub": user.Email,
//		"exp": time.Now().Add(time.Hour).Unix(),
//	})
//	tokenString, err := token.SignedString(privateKey)
//	if err != nil {
//		return "nil" //fmt.Sprintf("failed to creat token : %v", err)
//	} else {
//		return tokenString
//	}
//}
//type Credentials struct {
//	Email    string `json:"username"`
//	Password string `json:"password"`
//}
