package twosum

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for n := 1; n < len(nums); n++ {
			b := nums[n]
			if a+b == target {
				result = append(result, i)
				result = append(result, n)
			}
		}
	}

	return result
}
