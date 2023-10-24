// This the backend server for Study Schedular.
// It handles the routing and database access.
package main

import (
	"log"
	"net/http"
)

// Todo is the basic unit of the app. It is moved around for scheduling
type Todo struct {
	Title     string
	Urgent    bool
	Important bool
}

func home(_ http.ResponseWriter, _ *http.Request) {
	todos := []Todo{
		{Title: "Hindi Homework"},
		{Title: "English Worksheetk"},
		{Title: "Hindi 2 Q/A"},
	}

}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
