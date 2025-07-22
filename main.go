package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Courses struct{
	ID string `json:"id"`
	Title string `json:"name"`
}

var courses = []Courses{
	{"100","Grokking Modern System Design "},
	{"101","CloudLab: WebSockets-based Chat Application using API Gateway"},
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(courses)
	if err != nil{
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(jsonData))
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(courses)
}

func main() {
	http.HandleFunc("/courses", getHandler)
	http.ListenAndServe(":8080", nil)
}
