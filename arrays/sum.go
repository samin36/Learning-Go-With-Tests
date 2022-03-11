package arrays

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}

	return
}

func SumAll(numSlices ...[]int) (sums []int) {
	for _, nums := range numSlices {
		sums = append(sums, Sum(nums))
	}

	return
}

func SumAllTails(numSlices ...[]int) (tailSums []int) {
	for _, nums := range numSlices {
		tail := []int{}

		if len(nums) > 1 {
			tail = nums[1:]
		}

		tailSums = append(tailSums, Sum(tail))
	}

	return
}
