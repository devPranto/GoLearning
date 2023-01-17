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

var (
	mySigningKey = []byte("AllYourBase")
	expireTime   = time.Now().Add(2 * time.Hour)
)

// Login and Register serves the pages only , no logic used

func Login(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("views/login.html")
	t.Execute(writer, nil)

}

func Register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/register.html")
	t.Execute(w, nil)

}
func UpdateInfo(w http.ResponseWriter, r *http.Request) {
	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)

	user := models.User{
		Email:     r.FormValue("email"),
		FirstName: r.FormValue("fname"),
		LastName:  r.FormValue("lname"),
		Password:  password,
	}
	models.Update(&user)
	fmt.Println("Hi", r.FormValue("fname"))
	//t, _ := template.ParseFiles("views/register.html")
	//t.Execute(w, nil)
	w.Write([]byte("Success"))
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
		fmt.Println("hiiii : ", err.Error())
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
	user := models.Find(claims.Issuer)
	t, _ := template.ParseFiles("views/update.html")
	t.Execute(w, user)

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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
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
	str := strings.Replace(tempFile.Name(), "\\", "/", 1)
	userData.Path = str
	tempFile.Write(fileBytes)
	fmt.Println("==================================================================")
	fmt.Println("User Data : ", userData)
	fmt.Println("==================================================================")
	fmt.Fprint(w, userData)
	userData.Insert()

}

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
	t, err := template.ParseFiles("views/welcome.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result.JWT = token
	models.InsertToken(result)
	fmt.Println("In  Authe")
	t.Execute(writer, result)
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
	t, _ := template.ParseFiles("views/register.html")
	t.Execute(writer, nil)
	//todo remove token from mongodb
}
