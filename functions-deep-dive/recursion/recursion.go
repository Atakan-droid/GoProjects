package recursion

// Recursion
func main() {
	number := 5
	result := factorial(number)
	println(result)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
