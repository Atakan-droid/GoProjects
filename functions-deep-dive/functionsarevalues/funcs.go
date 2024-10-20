package functionsarevalues

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	transformer := getTransformerFunc(&numbers)

	doubledNumbers := transformNumbers(&numbers, transformer)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := make([]int, len(*numbers))
	for _, n := range *numbers {
		dNumbers = append(dNumbers, double(n))
	}
	return dNumbers
}

func getTransformerFunc(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := make([]int, len(*numbers))
	for _, n := range *numbers {
		dNumbers = append(dNumbers, transform(n))
	}
	return dNumbers
}
