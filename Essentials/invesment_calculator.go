package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")

	var res, _ = calculate()
	fmt.Println(res)

	// profit_calculator()
}

func calculate() (float64, error) {
	const inflitionRate float64 = 3.5
	var investmentAmount, years float64
	expectecReturnValue := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount) //Pointer to the variable

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectecReturnValue/100, years)
	futureRealValue := futureValue / math.Pow(1+inflitionRate/100, years)

	fmt.Println("Future Value: ", futureRealValue)
	return futureValue, nil
}
