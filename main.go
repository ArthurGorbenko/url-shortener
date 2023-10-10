package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/generate", controllers.CreateShortLink).Methods("POST")
	router.HandleFunc("/api/{hash}", controllers.GetOrigin).Methods("GET")

	controllers.SetupIndex(router)
}

func main() {
	LoadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}
