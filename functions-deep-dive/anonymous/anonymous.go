package anonymous

import "fmt"

// Anon function
func main() {
	numbers := []int{1, 2, 3, 4, 5}

	//TODO: Anonymous function

	double := createTransformer(2)

	doubled := transformNumbers(&numbers, func(number int) int { return number * 2 })

	transformed := transformNumbers(&numbers, double)

	fmt.Println(transformed)
	fmt.Println(doubled)
}

func double(number int) int {
	return number * 2
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := make([]int, len(*numbers))
	for _, n := range *numbers {
		dNumbers = append(dNumbers, transform(n))
	}
	return dNumbers
}

// TODO: Create a function that returns a function that multiplies a number by a factor
// TODO: factory function
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
