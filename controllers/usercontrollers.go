package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pbutarbutar/crud-golang/database"
	"github.com/pbutarbutar/crud-golang/entity"
	"github.com/pbutarbutar/crud-golang/utils"

	"github.com/gorilla/mux"
)

//GetAllUsers get all users data
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []entity.Users
	database.Connector.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

//GetUserByID returns user with specific UserID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["user_id"]

	var users entity.Users
	database.Connector.First(&users, "user_id = ? ", key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

//CreateUser creates user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.Users
	json.Unmarshal(requestBody, &user)

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		apiResp := entity.ApiResponse{}
		apiResp.Status = http.StatusBadRequest
		apiResp.Success = false
		apiResp.Message = "Erro Hashpassword"
		json.NewEncoder(w).Encode(apiResp)
		return
	}

	user.Password = password

	database.Connector.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

//UpdateUserByID updates user with respective ID
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.Users
	json.Unmarshal(requestBody, &user)

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		apiResp := entity.ApiResponse{}
		apiResp.Status = http.StatusBadRequest
		apiResp.Success = false
		apiResp.Message = "Erro Hashpassword"
		json.NewEncoder(w).Encode(apiResp)
		return
	}

	user.Password = password

	database.Connector.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//DeletUserByID delete's user with specific userId
func DeletUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["user_id"]

	var user entity.Users
	database.Connector.Where("user_id = ?", key).Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
