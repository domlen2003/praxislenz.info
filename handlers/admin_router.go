package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

var tpl *template.Template

var store = sessions.NewCookieStore([]byte("secret")) //TODO: replace with os.Getenv("SESSION_KEY")

func AddAdminRouter(r *mux.Router, t *template.Template) {
	tpl = t
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/loginauth", loginAuthHandler).Methods("POST")
	r.HandleFunc("/settings", settingsHandler)
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	fmt.Fprint(w, "Settings Page")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//redirect To Login Page
	http.Redirect(w, r, r.URL.Path+"/login", http.StatusPermanentRedirect)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//Show Form for Login
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("Username: ", username, " Password: ", password)

	if username != "Silvia Lenz" && password != "12345" { //TODO: replace with os.Getenv("PWD"), os.Getenv("USERNAME")
		tpl.ExecuteTemplate(w, "login.gohtml", "Username or Password wrong!")
	} else {
		session, _ := store.Get(r, "session")
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/admin/settings", http.StatusPermanentRedirect)
	}
}
