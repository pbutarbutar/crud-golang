package main

import (
	"log"
	"net/http"

	"github.com/pbutarbutar/crud-golang/controllers"
	"github.com/pbutarbutar/crud-golang/database"
	"github.com/pbutarbutar/crud-golang/entity"

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
	router.HandleFunc("/get", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletUserByID).Methods("DELETE")
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