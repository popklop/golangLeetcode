func maximumDifference(nums []int) int {
	minit := math.MaxInt
	result := math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < minit {
			minit = nums[i]
			for j := i; j < len(nums); j++ {
				if nums[j] > minit {
					if (nums[j] - minit) > result {
						result = nums[j] - minit
					}

				}
			}
		}
	}
	if result == math.MinInt {
		return -1
	}
	return result
}
