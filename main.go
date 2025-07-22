package main

import (
	"fmt"
	"net/http"
)



func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from SigmaWater")
}

func main() {
	http.HandleFunc("/", getHandler)
	http.ListenAndServe(":8080", nil)
}
