package main

//strucks are user-defined data types that can hold multiple values of different types
//they are similar to classes in other programming languages
//they are used to group related data together
//they are used to create complex data structures
//they are used to create reusable code

import (
	"advance_types/user"
	"fmt"
	"strconv"
	"time"
)

type customString string

func (customString) log() {
	fmt.Println("Logging")
}

func main() {
	//special method
	stuctFunc()

	var name customString = "John Doe"
	fmt.Println(name)
	name.log()
}

func stuctFunc() bool {
	var firstName, lastName string
	var birthDate time.Time
	var age int64

	firstName = getUserData("Enter First Name: ")
	lastName = getUserData("Enter Last Name: ")
	birthDate = time.Date(1976, time.January, 1, 0, 0, 0, 0, time.UTC)
	age, _ = strconv.ParseInt(getUserData("Enter Age: "), 10, 64)

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthDate, int(age))
	if err != nil {
		fmt.Println(err)
		return true
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	admin, err := user.NewAdmin("admin@gmail.com", "admin123")
	if err != nil {
		fmt.Println(err)
		return true
	}
	admin.OutputUserDetails()
	return false
}

func getUserData(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scanln(&input)
	return input
}
