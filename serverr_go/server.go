package main

import (
	"encoding/json"
	//"html/template"
	"net/http"
)

type Task struct {
	Name string
	Done bool
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/tasks", TaskHandler)
	http.ListenAndServe(":8888", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	task := Task{
		Name: "Server http",
		Done: true,
	}

	j, _ := json.Marshal(task)

	w.Write(j)

	//t := template.Must(template.ParseFiles("task.html"))
	//t.Execute(w, task)
}