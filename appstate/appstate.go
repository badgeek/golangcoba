package appstate

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

//AppState Struct
type AppState struct {
	Conn     *sql.DB
	DbUrl    string
	LogInfo  *log.Logger
	LogError *log.Logger
}

func (app *AppState) TestApp() {
	fmt.Println(app.DbUrl)
}

//App Instance
var App AppState

func SetupLog() {
	App.LogError = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	App.LogInfo = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}

func GetDBUrl() string {
	return App.DbUrl
}

func GetConnection() *sql.DB {
	return App.Conn
}

func SetupDB(url string) {
	App.DbUrl = url
}

func ConnectDB() {
	var err error
	App.Conn, err = sql.Open("postgres", GetDBUrl())
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func CloseDB() {
	App.Conn.Close()
}
