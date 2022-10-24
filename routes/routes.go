package routes

import (
	"github.com/gorilla/mux"
	"github.com/imajean/go-todo-postresql/controllers"
)

var RegisterTodoRoutes = func(router *mux.Router) {
	router.HandleFunc("/todos", controllers.ListTodos).Methods("GET")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
}
