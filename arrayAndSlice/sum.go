package arrayAndSlice

func Sum(arr []int) int {
	sum := 0

	for _, array := range arr {
		sum += array
	}

	return sum
}

func SumAll(numbsToSum ...[]int) []int {

	var sums []int
	for _, numbs := range numbsToSum {
		sums = append(sums, Sum(numbs))
	}

	return sums
}

func SumAllTails(numbsToSumTails ...[]int) []int {

	var sums []int
	for _, numbs := range numbsToSumTails {
		if len(numbs) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbs[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
