package main

import "fmt"

type User struct {
	Name string
	Age int
}

// by value (does Not persist)
func updateAgeByValue(u User, newAge int) {
	u.Age = newAge
}

// by pointer (persists)
func updateAgeByPointer(u *User, newAge int) {
	u.Age = newAge
}

func main() {
	user := User{Name: "Henry", Age: 21}

	updateAgeByValue(user, 25)
	fmt.Println("After updateAgeByValue:", user.Age)

	updateAgeByPointer(&user, 25)
	fmt.Println("After updateAgeByPointer:", user.Age)
}