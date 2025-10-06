package main

import (
	"html/template"
	"log"
	"net/http"
)

type Student struct {
	Name  string
	Age   string
	Quote string
	Hobby string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func IdCard(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	quote := r.URL.Query().Get("quote")
	hobby := r.URL.Query().Get("hobby")
	student := Student{
		Name:  name,
		Age:   age,
		Quote: quote,
		Hobby: hobby,
	}
	tmpl, err := template.ParseFiles("Page/idCard.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, student)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/idcard", IdCard)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
