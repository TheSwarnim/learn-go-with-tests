package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	// new refactored code
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum) // refactor
	// sums := make([]int, lengthOfNumbers)

	var sums []int

	// for i, numbers := range numbersToSum {
	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
