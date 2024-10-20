package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, -5}
	sum := sumup(1, 2, 3, 4, -5)
	sumOther := sumup(1, numbers[1:]...)
	fmt.Println(sum)
	fmt.Println(sumOther)
}

// Variadic function
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, n := range numbers {
		sum += n
	}
	return sum
}
