package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"html/template"
	"log"
	"net/http"
	"praxislenz.info/handlers"
	"praxislenz.info/middleware"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	//Routing
	r := mux.NewRouter()
	handlers.AddAdminRouter(r.PathPrefix("/admin").Subrouter(), tpl)
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	//Middleware
	errorChain := alice.New(middleware.LoggerHandler, middleware.RecoverHandler)
	http.Handle("/", errorChain.Then(r))
	http.Handle("/assets/", errorChain.Then(http.StripPrefix("/assets", http.FileServer(http.Dir("./templates/assets")))))

	//Server
	//err :=http.Serve(autocert.NewListener("domain.com"), nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server: ", err)
	}
}

type IndexContent struct {
	Title   string
	Message string
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	/*	tmpl := template.Must(template.ParseFiles("praxislenz.info/templates/index.gohtml"))
		data := IndexContent{
			Title: "Hello",
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Fatal("IndexHandler: ", err)
		}*/
}
