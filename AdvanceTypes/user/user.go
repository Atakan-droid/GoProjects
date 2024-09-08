package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthDate   time.Time
	age         int
	createdDate time.Time
}

type Admin struct {
	email    string
	password string
	User     //embedded struct
}

func NewAdmin(email, password string) (Admin, error) {
	if email == "" || password == "" {
		return Admin{}, errors.New("Invalid Admin Data")
	}
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName:   "Admin",
			lastName:    "User",
			birthDate:   time.Date(1976, time.January, 1, 0, 0, 0, 0, time.UTC),
			age:         44,
			createdDate: time.Now(),
		},
	}, nil
}

// access the struct using a pointer
func (u User) OutputUserDetails() {
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("Birth Date: ", u.birthDate)
	fmt.Println("Age: ", u.age)
	fmt.Println("Created Date: ", u.createdDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName string, birthDate time.Time, age int) (*User, error) {
	//create a new user
	if firstName == "" || lastName == "" || age < 0 {
		return nil, errors.New("Invalid User Data")
	}
	return &User{
		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthDate,
		age:         int(age),
		createdDate: time.Now(),
	}, nil
}
