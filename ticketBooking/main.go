package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct{
	Title string
	Body []byte
}

func handler(w http.ResponseWriter , r *http.Request){
	// fmt.Fprintf(w,"<h1>Hi this is dev %s ! </h1>",r.URL.Path[1:])
	page := &Page{}
	renderTemplate(w,"edit",page)
}

func viewHandler(w http.ResponseWriter , r *http.Request){
	title := r.URL.Path[len("/view/"):]
	page,err := load(title+".txt")
	fmt.Println(page)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(w,"view",page)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := load(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter,r *http.Request){
	//title := r.URL.Path[len("/save/"):]
	title := r.FormValue("title")
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()   
	fmt.Println(p)  
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	 t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func (p *Page)save()error{
	filename:= p.Title+".txt"
	return os.WriteFile(filename,p.Body,0600)
}

func load(title string) (*Page,error){
	filename := title+".txt"
	body, err := os.ReadFile(filename)
	if err != nil{
		return nil,err
	}
	return &Page{Title: title, Body: body},nil
}

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/view/",viewHandler)
	http.HandleFunc("/edit/",editHandler)
	http.HandleFunc("/save/",saveHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))

}
