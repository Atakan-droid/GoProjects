package main

import "fmt"

//global variables are defined outside of the main function
const someText = "Hello World"

var accountBalance float64 = 1000.0

func main() {
	//formatting floats
	// formattingFloats()

	//understanding functions
	// returned := understandingFunctions("Hello", "World")
	// fmt.Println(returned)

	//calculated future value
	// futureValue := calculatedFutureValue()
	// fmt.Println(futureValue)

	//control structures
	//if else
	//controlStructures(10)

	//bank
	//controlBank()

	//loops
	controlLoops()
}

func controlLoops() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			//break out of the loop
			break
		}
	}
}

func controlBank() {
	fmt.Println("Welcome to the bank")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Deposit")
	case 2:
		if accountBalance <= 0 {
			fmt.Println("Insufficient funds")
			return
		}
		fmt.Print("Enter amount to withdraw: ")
		var amount float64
		fmt.Scanln(&amount)
		accountBalance -= amount
		fmt.Println("Account balance is: ", accountBalance)
	case 3:
		fmt.Println("Thank you for visiting the bank")
		return
	default:
		fmt.Println("Invalid choice, please try again")
		controlBank()
	}

	controlBank()
}

func controlStructures(i int) {
	if i > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Less than 10")
	}

	//switch case
	switch i {
	case 10:
		fmt.Println("Equal to 10")
	case 11:
		fmt.Println("Equal to 11")
	default:
		fmt.Println("Not equal to 10 or 11")
	}
}

//function inputs are defined by the type of the variable
//function outputs are defined by the type of the return value
func understandingFunctions(text string, text2 string) string {
	fmt.Println(text)
	fmt.Println(text2)
	fmt.Println(someText)
	return "Done"
}

func calculatedFutureValue() (fv float64) {
	//calculating future value
	//present value * (1 + interest rate) ^ years
	presentValue := 1000.0
	interestRate := 0.05
	years := 5
	futureValue := float64(presentValue) * (1 + interestRate) * float64(years)
	return futureValue
}

func formattingFloats() {
	// %f is the default format, which is 6 decimal places
	fmt.Printf("%.2f\n", 3.14159265359)
	fmt.Printf("%.3f\n", 3.14159265359)

	//crating formatted strings
	fmt.Printf("%s is a %s\n", "This", "string")
	fmt.Printf("%s is a %d\n", "This", 1)
	fmt.Printf("%s is a %t\n", "This", true)
	fmt.Printf("%s is a %v\n", "This", 1.2)

	//building multi-line strings
	fmt.Printf(`This is a multi-line`)
}
