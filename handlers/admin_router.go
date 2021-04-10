package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var tpl *template.Template

func AddAdminRouter(r *mux.Router, t *template.Template) {
	tpl = t
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/loginauth", loginAuthHandler).Methods("POST")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	/*	_, err := fmt.Fprintln(w, "Admin")
		if err != nil {
			log.Fatal("IndexHandler: ", err)
		}*/
	//redirect To Login Page

	http.Redirect(w, r, r.URL.Path+"/login", http.StatusPermanentRedirect)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("Username: ", username, " Password: ", password)

	if username != "Silvia Lenz" && password != "12345" {
		tpl.ExecuteTemplate(w, "login.gohtml", "Username or Password wrong!")
	} else {
		fmt.Fprint(w, "You´re logged in successfully.")
	}
}
