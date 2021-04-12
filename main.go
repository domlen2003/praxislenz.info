package main

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/justinas/alice"
	"html/template"
	"log"
	"net/http"
	"praxislenz.info/handlers"
	"praxislenz.info/middleware"
	"time"
)

var tpl *template.Template

func main() {
	//ENV
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//MONGO
	handlers.StartMongoHandler()

	//Gohtml Templates
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	//Routing
	r := mux.NewRouter()
	handlers.AddAdminRouter(r.PathPrefix("/admin").Subrouter(), tpl)
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	//Middleware and File-Server
	errorChain := alice.New(middleware.LoggerHandler, middleware.RecoverHandler)
	http.Handle("/", errorChain.Then(r))
	http.Handle("/assets/", errorChain.Then(http.StripPrefix("/assets", http.FileServer(http.Dir("./templates/assets")))))

	// serve HTTPS
	http.ListenAndServe(":80", nil)
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Minute, // 5 min to allow for delays when 'curl' on OSx prompts for username/password
		WriteTimeout: 10 * time.Second,
		TLSConfig:    &tls.Config{ServerName: "praxislenz.info"},
	}
	err = server.ListenAndServeTLS(".pem", ".key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type IndexContent struct {
	CoronaContent   string
	CoronaTimestamp string
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	cinfo := handlers.GetInfo(handlers.CoronaInfo)[0]
	data := IndexContent{cinfo.Content, cinfo.Timestamp}
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatal("Index: ", err)
	}
}
