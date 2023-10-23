// This the backend server for Study Schedular.
// It handles the routing and database access.
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Todo is the basic unit of the app. It is moved around for scheduling
type Todo struct {
	Title string
}

func home(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/index.html"))
	todos := map[string][]Todo{
		"Todos": {
			{Title: "Hindi Homework"},
			{Title: "English Worksheetk"},
			{Title: "Hindi 2 Q/A"},
		},
	}

	tmpl.Execute(w, todos)
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
