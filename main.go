package main

import (
	"log"
	"net/http"

	"github.com/pbutarbutar/crud-golang/controllers"
	"github.com/pbutarbutar/crud-golang/database"
	"github.com/pbutarbutar/crud-golang/entity"
	"github.com/pbutarbutar/crud-golang/utils"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Test Parul - Starting server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	router.Handle("/get", utils.MiddlewareValidateToken(http.HandlerFunc(controllers.GetAllUsers))).Methods("GET")
	router.Handle("/get/{user_id}", utils.MiddlewareValidateToken(http.HandlerFunc(controllers.GetUserByID))).Methods("GET")
	router.Handle("/update/{user_id}", utils.MiddlewareValidateToken(http.HandlerFunc(controllers.UpdateUserByID))).Methods("PUT")
	router.Handle("/delete/{user_id}", utils.MiddlewareValidateToken(http.HandlerFunc(controllers.DeletUserByID))).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "qwerty123",
			DB:         "gotest",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Users{})
}
