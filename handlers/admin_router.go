package handlers

import (
	"fmt"
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

func AddAdminRouter(r *mux.Router, t *template.Template) {
	tpl = t
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/loginauth", loginAuthHandler).Methods("POST")
	r.HandleFunc("/settings", settingsHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//redirect To Login Page
	http.Redirect(w, r, r.URL.Path+"/login", http.StatusPermanentRedirect)
}
func loginHandler(w http.ResponseWriter, _ *http.Request) {
	//Show Form for Login
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("Username: ", username, " Password: ", password)

	if username != os.Getenv("ADMIN_USERNAME") && password != os.Getenv("ADMIN_PASSWORD") {
		tpl.ExecuteTemplate(w, "login.gohtml", "Username or Password wrong!")
	} else {
		session, _ := store.Get(r, "session")
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/admin/settings", http.StatusPermanentRedirect)
	}
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		r.ParseForm()
		if len(r.FormValue("contentType")) > 0 && len(r.FormValue("content")) > 0 {
			fmt.Println("Typ: " + r.FormValue("contentType"))
			values := strings.Split(strings.ReplaceAll(r.FormValue("content"), "\r\n", "\n"), "\n")
			fmt.Printf("Values: %v\n", values)
			fmt.Println(time.Now().Format("2.1.2006 15:04"))
		}
		tpl.ExecuteTemplate(w, "settings.gohtml", indexSettings{ContentTypes: contentTypes, CurrentContent: "Current Content"})
	}
}

type indexSettings struct {
	ContentTypes   []string
	CurrentContent string
}

var contentTypes = []string{"Corona Info", "Generelle Info"}
