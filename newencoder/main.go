// encoding json data using the newEncoder function

package main

import (
	"net/http"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	person := Person{ Name: "John", Age: 30}

	// encoding 
	encoder := json.NewEncoder(w)
	err := encoder.Encode(person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}