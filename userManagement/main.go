package main

// using pointes in a simple user management 

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// a simple user model
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// pretend db thats in-memory
var db = map[int]*User {
	1: {ID: 1, Name: "Henry", Email: "henryarnoldme@gmail.com"},
}

// update function uses *User so that changes persist
func updateUser(u *User, newName string, newEmail string) {
	u.Name = newName
	u.Email = newEmail
}

// http handler
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := db[1] // pretends to fetch user with id:1

	// parse request body
	var payload struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	// update in place (pointer ensures we change the original object in the db)
	updateUser(user, payload.Name, payload.Email)

	// responds with updated user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/update", updateUserHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}