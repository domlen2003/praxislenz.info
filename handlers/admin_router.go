package handlers

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// AddAdminRouter define the routes within the /admin subrouter and loads the templates
func AddAdminRouter(r *mux.Router, t *template.Template) {
	tpl = t
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/loginauth", loginAuthHandler).Methods("POST")
	r.HandleFunc("/settings", settingsHandler)
}

//indexHandler always redirects to /login since there is no way to sign up etc. at the moment
func indexHandler(w http.ResponseWriter, r *http.Request) {
	//redirect To Login Page
	http.Redirect(w, r, r.URL.Path+"/login", http.StatusPermanentRedirect)
}

//loginHandler initially renders the login form so that /loginouth only has to listen to the more secure POST
func loginHandler(w http.ResponseWriter, _ *http.Request) {
	//Show Form for Login
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

//loginAuthHandler compares the given username and password with the .env and redirects accordingly
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != os.Getenv("ADMIN_USERNAME") && password != os.Getenv("ADMIN_PASSWORD") {
		tpl.ExecuteTemplate(w, "login.gohtml", "Username or Password wrong!")
	} else {
		session, _ := store.Get(r, "session")
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/admin/settings", http.StatusPermanentRedirect)
	}
}

//settingsHandler renders the form for updating the indexes content and updates mongodb
func settingsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		r.ParseForm()
		if len(r.FormValue("contentType")) > 0 && len(r.FormValue("content")) > 0 {
			replacer := strings.NewReplacer("\r\n", "<br>", "\n", "<br>")
			UpdateInfo(InfoNode{
				Type:      Infotype(r.FormValue("contentType")),
				Content:   replacer.Replace(r.FormValue("content")),
				Timestamp: time.Now().Format("2.1.2006 15:04"),
			})
		}
		tpl.ExecuteTemplate(w, "settings.gohtml", indexSettings{ContentTypes: contentTypes})
	}
}

type indexSettings struct {
	ContentTypes []string
}

var contentTypes = []string{string(CoronaInfo), string(GeneralInfo), string(OpeningHours)}
