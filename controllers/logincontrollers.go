package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pbutarbutar/crud-golang/database"
	"github.com/pbutarbutar/crud-golang/entity"
	"github.com/pbutarbutar/crud-golang/utils"
)

//LoginUser is struct login request
type LoginUser struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

//Login is endpoint login
func Login(w http.ResponseWriter, r *http.Request) {

	apiResp := entity.ApiResponse{}

	requestBody, _ := ioutil.ReadAll(r.Body)

	var login LoginUser
	json.Unmarshal(requestBody, &login)

	var users entity.Users
	if err := database.Connector.First(&users, "user_id = ? ", login.UserID).Error; err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		apiResp.Status = http.StatusUnauthorized
		apiResp.Success = false
		apiResp.Message = "Login Failed - Your userid not found!"
		json.NewEncoder(w).Encode(apiResp)
		return
	}

	isLogin := utils.CheckPasswordHash(login.Password, users.Password)

	if !isLogin {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		apiResp.Status = http.StatusUnauthorized
		apiResp.Success = false
		apiResp.Message = "Login Failed - Your password not found"
		json.NewEncoder(w).Encode(apiResp)
		return
	}

	token, err := utils.GetAuthToken(&users)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		apiResp.Status = http.StatusUnauthorized
		apiResp.Success = false
		apiResp.Message = "Generate Token Failed"
		json.NewEncoder(w).Encode(apiResp)
		return
	}

	apiResp.Status = 200
	apiResp.Success = true
	apiResp.Message = "Login Success"
	apiResp.Data = map[string]interface{}{
		"UserID": users.UserID,
		"token":  token,
		"roleID": "-",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResp)
}
