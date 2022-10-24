package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imajean/go-todo-postresql/database"
	"github.com/imajean/go-todo-postresql/models"
	"gorm.io/gorm"
)

func db() *gorm.DB {
	return database.GetDB()
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListTodos")
	var todos []models.Todo
	db().Find(&todos)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateTodo")
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	fmt.Println("TODO", todo)
	createdTodo := db().Create(&todo)

	w.Header().Set("Content-Type", "application/json")
	if createdTodo.Error != nil {
		json.NewEncoder(w).Encode(createdTodo.Error)
	} else {
		json.NewEncoder(w).Encode(&todo)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetTodo")
	params := mux.Vars(r)
	id, _ := (strconv.ParseInt(params["id"], 10, 32))
	var todo models.Todo
	result := db().First(&todo, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Todo not found : %s", params["id"])
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateTodo")
	params := mux.Vars(r)
	id, _ := (strconv.ParseInt(params["id"], 10, 32))
	var todo models.Todo
	result := db().First(&todo, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Todo not found : %s", params["id"])
		return
	}

	var updatedTodo models.Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)

	todo.Title = updatedTodo.Title
	todo.Done = updatedTodo.Done

	fmt.Println("TODO", todo)
	db().Save(&todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteTodo")
	params := mux.Vars(r)

	var todo models.Todo

	result := db().First(&todo, params["id"])

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Todo not found : %s", params["id"])
		return
	}

	db().Delete(&todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
