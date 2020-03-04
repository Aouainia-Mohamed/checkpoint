package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", loggg)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)

}

func loggg(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	email := r.FormValue("mail")
	password := r.FormValue("pwd")

	d := struct {
		Eemail    string
		Ppassword string
	}{
		Eemail:    email,
		Ppassword: password,
	}

	tpl.ExecuteTemplate(w, "login.gohtml", d)

}
