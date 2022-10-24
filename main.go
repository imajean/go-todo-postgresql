package main

import (
	"log"
	"os"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/imajean/go-todo-postresql/database"
	"github.com/imajean/go-todo-postresql/routes"
	"google.golang.org/appengine"
)

// fix issue where routes /todos and /todos/ does not go to the same route handlerFunc
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

func main() {
	database.InitDbConnection()

	// API routes
	router := mux.NewRouter()
	routes.RegisterTodoRoutes(router)

	// Called when the instance is started by App Engine
	http.HandleFunc("/_ah/start", func(w http.ResponseWriter, r *http.Request) {
	})
	// Called when the instance is stopped by App Engine
	http.HandleFunc("/_ah/stop", func(w http.ResponseWriter, r *http.Request) {
	})

	http.Handle("/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, removeTrailingSlash(router)); err != nil {
		log.Fatal(err)
	}

	appengine.Main() // Start the server
}
