func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	i := 0
	j := (n + 1) / 2

	for j < n {
		if nums[i] < nums[j] {
			i++ 
		}
		j++ 
	}
	return n - 2*i 
}