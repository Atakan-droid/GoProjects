package main

import (
	"fmt"
	"packages_module/calculator"
)

func main() {
	res := calculator.Calculate(10, 5, "+")
	fmt.Println(res)
}
