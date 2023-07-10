package controllers

import (
	"basic_api/config"
	"basic_api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	db := config.Dbmigration()
	params := mux.Vars(r)
	inputtodoid := params["todoid"]

	var todo models.Todo
	db.Preload("Catgories").First(&todo, inputtodoid)
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var todos []models.Todo
	db := config.Dbmigration()
	db.Preload("Catgories").Find(&todos)
	json.NewEncoder(w).Encode(todos)
}
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var todo models.Todo
	db := config.Dbmigration()
	json.NewDecoder(r.Body).Decode(&todo)
	db.Create(&todo)
	json.NewEncoder(w).Encode("todo created")
	json.NewEncoder(w).Encode(todo)
}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var todo models.Todo
	db := config.Dbmigration()
	db.First(&todo, params["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	db.Save(todo)
	json.NewEncoder(w).Encode(todo)

}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	db := config.Dbmigration()
	params := mux.Vars(r)
	todoid := params["todoid"]
	// Convert `orderId` string param to uint64
	id64, _ := strconv.ParseUint(todoid, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)
	db.Where("todo_id = ?", idToDelete).Delete(&models.Todo{})
	db.Where("todo_id = ?", idToDelete).Delete(&models.Category{})
	json.NewEncoder(w).Encode("todo is deleted")
}
