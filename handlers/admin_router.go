package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func AddAdminRouter(r *mux.Router) {
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/test", Test)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	/*	_, err := fmt.Fprintln(w, "Admin")
		if err != nil {
			log.Fatal("IndexHandler: ", err)
		}*/
	//redirect To Login Page

	http.Redirect(w, r, r.URL.Path+"/login", http.StatusPermanentRedirect)
}

func loginHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "Login")
	if err != nil {
		log.Fatal("IndexHandler: ", err)
	}
}

type HelloResponse struct {
	Message string `json:"message"`
}

type TodoPageData struct {
	Fontsize  int
	PageTitle string
	Todos     []Todo
}

type Todo struct {
	Title string
	Done  bool
}

func Test(w http.ResponseWriter, _ *http.Request) {

	tmpl := template.Must(template.ParseFiles("views/layout.html"))

	data := TodoPageData{
		Fontsize:  20,
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal("IndexHandler: ", err)
	}
}
