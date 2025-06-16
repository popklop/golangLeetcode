func buildArray(nums []int) []int {
	arr:=make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[i] = nums[nums[i]]

	}

	return arr
}
