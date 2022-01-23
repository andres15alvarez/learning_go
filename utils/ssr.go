package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var todos = []Todo{
	{Title: "Task 1", IsDone: false},
	{Title: "Task 2", IsDone: true},
	{Title: "Task 3", IsDone: true},
}

func RunServer() {
	tmpl := template.Must(template.ParseFiles("utils/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     todos,
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(todos)
	})

	fmt.Print("Server running in port: 8000...")
	http.ListenAndServe(":8000", nil)

}
