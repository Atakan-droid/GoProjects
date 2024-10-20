package main

import "fmt"

func practice() {

	//Create a new Array of hobbies
	hobbies := [5]string{"Programming", "Gaming", "Reading"}
	fmt.Println(hobbies)

	//Also add more data to the array
	hobbies[3] = "Swimming"
	hobbies[4] = "Cycling"
	fmt.Println(hobbies)

	//Create a slice of the array
	featuredHobbies := hobbies[:2]
	fmt.Println(featuredHobbies)

	//Re-slice the slice from 3 to 5
	mainHobbies := featuredHobbies[1:]
	fmt.Println(mainHobbies)

	//Create a dynamic array course goals
	courseGoals := []string{"Learn Go", "Build a project"}
	fmt.Println(courseGoals)

	//Set the second goal and different goal
	courseGoals[1] = "Build a web app"
	courseGoals = append(courseGoals, "Get a job")
	fmt.Println(courseGoals)

	//Create a Product struct and title id,rpice and
	//create dynamic list at least 2 products
	//then add third proudtc to esixting list
	type Product struct {
		id    int
		name  string
		price float64
	}

	products := []Product{
		{1, "Apple", 1.2},
		{2, "Banana", 3.4},
	}
	fmt.Println(products)

	products = append(products, Product{3, "Cherry", 5.6})
	fmt.Println(products)
}
