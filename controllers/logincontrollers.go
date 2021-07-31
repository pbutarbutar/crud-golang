package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pbutarbutar/crud-golang/database"
	"github.com/pbutarbutar/crud-golang/entity"

	"github.com/gorilla/mux"
)

//Login User
func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["user_id"]

	var users entity.Users
	database.Connector.First(&users, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
