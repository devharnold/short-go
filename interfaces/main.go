package main

import "fmt"

// define an interface
type Notifier interface {
	Notify() string
}

// user struct
type User struct {
	Name string
	Email string
}

// admin struct
type Admin struct {
	Name string
	Role string
}

// implement Notify() for user
func (u User) Notify() string {
	return fmt.Sprintf("User %s notified via email: %s", u.Name, u.Email)
}

// implement Notify() for admin
func (a Admin) Notify() string {
	return fmt.Sprintf("Admin %s notified via email: %s", a.Name, a.Role)
}

// function that works with any notifier
func sendNotification(n Notifier) {
	fmt.Println(n.Notify())
}

func main() {
	user := User{Name: "Henry", Email: "henryarnoldme@gmail.com"}
	admin := Admin{Name: "Admin", Role: "SuperAdmin"}

	sendNotification(user)
	sendNotification(admin)
}

