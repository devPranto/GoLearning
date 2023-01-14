package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"login-project/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("views/login.html")
	t.Execute(writer, nil)

}

func Authenticate(writer http.ResponseWriter, request *http.Request) {

	Email := strings.TrimSpace(request.FormValue("email"))
	password := strings.TrimSpace(request.FormValue("password"))
	result := models.Find(Email)
	fmt.Println(result)
	err := bcrypt.CompareHashAndPassword(result.Password, []byte(password))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "bad request")
		return
	}
	token := CreatJWT(result)
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
func Update(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("JWT")
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwt.SigningMethodES256, nil
		})
	fmt.Println(claims.Issuer)
	//models.Find(id)
	//fmt.Printf("Claim : %v \n", claims)
	//for i, v := range claims {
	//	fmt.Printf("claim : %+v , %+v \n", i, v)
	//}
	fmt.Println(cookie)
	fmt.Println(token.Valid)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "unauthorized access")
		return
	}
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

func Register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/register.html")
	t.Execute(w, nil)

}

func CreatJWT(user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(token)
	if err != nil {
		return fmt.Sprintf("failed to creat token : %v", err)
	} else {
		return tokenString
	}
}

//"failed to creat token : key is of invalid type: expected *ecdsa.PrivateKey or crypto.Signer, received string"
