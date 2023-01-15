package controllers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io/ioutil"
	"log"
	"login-project/models"
	"net/http"
	"strings"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("views/login.html")
	t.Execute(writer, nil)

}

var (
	mySigningKey = []byte("AllYourBase")
	expireTime   = time.Now().Add(2 * time.Minute)
)

func Authenticate(writer http.ResponseWriter, request *http.Request) {

	Email := strings.TrimSpace(request.FormValue("email"))
	password := strings.TrimSpace(request.FormValue("password"))
	result := models.Find(Email)
	fmt.Printf(" Query Data: %v \n", result)
	err := bcrypt.CompareHashAndPassword(result.Password, []byte(password))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "bad request")
		return
	}
	token := CreatJWT(result)
	fmt.Printf("This is the token : %v \n", token)
	cookie := http.Cookie{
		Name:     "JWT",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(writer, &cookie)
	t, _ := template.ParseFiles("views/welcome.html")
	result.JWT = token
	models.InsertToken(result)
	t.Execute(writer, result)
}

func CreatJWT(user *models.User) string {

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		Issuer:    user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	return ss
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HI im update")
	cookie, err := r.Cookie("JWT")
	if err != nil {
		fmt.Println(err.Error())
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString := cookie.Value
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("hiiii")
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		fmt.Println("hiiii 2")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims)))

}
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		fmt.Println("HI")
		h.ServeHTTP(w, r)
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, "")

}

func Submit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseMultipartForm(10 << 20)
	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)
	userData := models.User{
		FirstName: strings.TrimSpace(r.FormValue("Fname")),
		LastName:  strings.TrimSpace(r.FormValue("Lname")),
		Email:     strings.TrimSpace(r.FormValue("email")),
		Gender:    r.FormValue("gender"),
		Password:  password,
	}
	//reading image
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fileName := "profile_" + userData.FirstName + "*.jpg"
	tempFile, err := ioutil.TempFile("./picture", fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	userData.Path = tempFile.Name()
	tempFile.Write(fileBytes)
	fmt.Println(userData)
	fmt.Fprint(w, userData)
	userData.Insert()

}
func Logout(writer http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "JWT",
		Value:    "",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(writer, &cookie)
	//todo remove token from mongodb
}
func Register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/register.html")
	t.Execute(w, nil)

}
