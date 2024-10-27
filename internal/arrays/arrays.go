package arrays

func Sum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

func SumAll(arrays ...[]int) []int {
	var sums []int

	for _, nums := range arrays {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(arrays ...[]int) []int {
	var tailsums []int

	for _, nums := range arrays {
		if len(nums) == 0 {
			tailsums = append(tailsums, 0)
			continue
		}
		tailsums = append(tailsums, Sum(nums[1:]))
	}

	return tailsums
}
