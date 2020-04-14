package math

// Sum will add a slice of integers and return the total
func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}