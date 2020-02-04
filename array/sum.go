package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(numbers ...[]int) (sums []int) {

	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}

	return
}
