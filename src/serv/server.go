package main

import (
	"net/http"
	"io/ioutil"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
	// read form value
	value := r.FormValue("value")
	if r.Method == "POST" {
		// receive posted data
		body, err := ioutil.ReadAll(r.Body)
	}
}
func main() {
	http.HandleFunc("/", someHandler)
	http.ListenAndServe(":8080", nil)
}
