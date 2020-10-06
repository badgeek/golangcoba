package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"

	appRepository "manticore.id/golangcoba/repository"
)

func Home(w http.ResponseWriter, r *http.Request) {

	templateFiles := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// greet := []byte("Hello VOID")
	if err != nil {
		w.Write([]byte("Invalid"))
	}

	if id > 0 {
		// w.Write([]byte("OK"))
		fmt.Fprintf(w, "Found id %d", id)
	} else {
		// w.Write([]byte("MEH"))
		http.NotFound(w, r)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	err := appRepository.CreateVersion(r.PostFormValue("version"))

	if err != nil {
		w.Write([]byte("Error writing new version"))
	} else {
		w.Write([]byte("Succesfuly add new version"))
	}
}

func ShowVersions(w http.ResponseWriter, r *http.Request) {
	// appRepository.GetVersions()

	verID := r.URL.Query().Get("id")

	res, _ := appRepository.IsVersionExist(verID)

	if res == true {
		w.Write([]byte("Versions OK"))
	} else {
		w.Write([]byte("Versions Not OK"))
	}
}

func ListVersion(w http.ResponseWriter, r *http.Request) {

	res, _ := appRepository.GetVersions()
	jsonResponse, err := json.Marshal(res)

	if err != nil {
		w.Write([]byte("Error Decoding JSON"))
	} else {
		w.Write(jsonResponse)
	}
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func ListRate(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
	w.Write(bodyBytes)
}
