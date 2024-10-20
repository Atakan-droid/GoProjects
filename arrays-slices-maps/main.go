package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) print() {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func main() {

	//Make function call
	//Make function is used to create a slice with a length and capacity
	userNames := make([]string, 0, 5)

	userNames = append(userNames, "John")
	userNames = append(userNames, "Doe")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Doe")
	userNames = append(userNames, "Doe")
	userNames = append(userNames, "Doe")
	userNames = append(userNames, "Doe")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)

	courseRatings["Java"] = 4.2
	courseRatings["Go"] = 4.7
	courseRatings["Python"] = 4.5

	fmt.Println(courseRatings)

	//Working with aliases
	//An alias is a new name given to an existing type

	courseRatings2 := make(floatMap, 3)

	courseRatings2["Java"] = 4.2
	courseRatings2["Go"] = 4.7
	courseRatings2["Python"] = 4.5

	courseRatings2.print()

	//for loop
	for _, username := range userNames {
		fmt.Println("Username: ", username)
	}

	for key, value := range courseRatings2 {
		fmt.Println("Key: ", key, "Value: ", value)
	}

}
