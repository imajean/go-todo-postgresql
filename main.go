package main

import (
	"fmt"
	"log"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/imajean/go-todo-postresql/database"
	"github.com/imajean/go-todo-postresql/routes"
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
	http.Handle("/", router)

	fmt.Println("Serving at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", removeTrailingSlash(router)))
}
