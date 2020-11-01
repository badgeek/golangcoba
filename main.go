package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	appstate "manticore.id/golangcoba/appstate"
	controller "manticore.id/golangcoba/controller"
)

// var app appstate.AppState

func setupRouting() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/versions", controller.ShowVersions)
	mux.HandleFunc("/versions/create", controller.Create)
	mux.HandleFunc("/versions/list", controller.ListVersion)
	mux.HandleFunc("/versions/listrate", controller.ListRate)
	return mux
}

func setupHTTP(route *http.ServeMux) {
	err := http.ListenAndServe(":8000", route)
	log.Fatal(err)
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
}

func main() {
	// Open handle to database like normal
	loadENV()
	appstate.SetupLog()
	appstate.SetupDB(os.Getenv("DBCONFIG"))
	appstate.ConnectDB()
	appstate.App.LogInfo.Printf("SERVER: Started")
	defer appstate.CloseDB()
	route := setupRouting()
	setupHTTP(route)
}
