package main

/*
All values in Go have a so-called "Null Value" - i.e., the value that's set as a default
if no value is assigned to a variable.

For example, the null value of an int variable is 0. Of a float64,
it would be 0.0. Of a string, it's "".

For a pointer, it's nil - a special value built-into Go.

nil represents the absence of an address value - i.e.,
a pointer pointing at no address / no value in memory.
*/

import "fmt"

func main() {
	age := 44 //regular variable

	agePointer := &age //pointer to the memory address of the variable

	fmt.Println("Address Of Age: ", agePointer)   //memory address of the variable
	fmt.Println("Value Of Age: ", *agePointer)    //value of the variable
	fmt.Println(getAdultYears(age))               //passing a variable
	fmt.Println(getAdultYearsPointer(agePointer)) //passing a pointer to a variable
}

func getAdultYears(age int) int {
	//passing a variable creates a copy of the variable so different memory address
	fmt.Println(&age) //pointer to the memory address of the variable
	return age - 18
}

func getAdultYearsPointer(agePointer *int) int {
	//passing a pointer to a variable so same memory address
	fmt.Println(agePointer) //memory address of the variable
	age := *agePointer
	return age - 18
}
