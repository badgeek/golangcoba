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

func setupHTTP() {
	mux := http.NewServeMux()
	mux.HandleFunc("/versions", controller.ShowVersions)
	mux.HandleFunc("/versions/create", controller.Create)
	mux.HandleFunc("/versions/list", controller.ListVersion)
	mux.HandleFunc("/versions/listrate", controller.ListRate)
	err := http.ListenAndServe(":8000", mux)
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

	// appstate.App.LogInfo = new log.l

	appstate.App.LogInfo.Printf("coba coba")

	defer appstate.CloseDB()
	setupHTTP()

	// c, err := semver.NewConstraint("<= 1.2.3, >= 1.4")
	// if err != nil {
	// 	// Handle constraint not being parseable.
	// }

	// v, err := semver.NewVersion("1.3")
	// if err != nil {
	// 	// Handle version not being parseable.
	// }

	// incp := v.IncMinor()
	// fmt.Println(incp)

	// // Validate a version against a constraint.
	// _, msgs := c.Validate(v)
	// // a is false
	// for _, m := range msgs {
	// 	fmt.Println(m)

	// 	// Loops over the errors which would read
	// 	// "1.3 is greater than 1.2.3"
	// 	// "1.3 is less than 1.4"
	// }
}
