// encoding json data using the Marshal function

package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	person := Person{ Name: "John", Age: 30}
	// encoding
	jsonStr, err := json.Marshal(person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonStr)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}