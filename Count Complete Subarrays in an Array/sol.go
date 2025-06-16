func countCompleteSubarrays(nums []int) int {
	mapa := make(map[int]int)
	for _, n := range nums {
		mapa[n] = 1
	}
	start := 0
	unique := 0
	newmapa := make(map[int]int)
	result := 0
	for end := 0; end < len(nums); end++ {
		if newmapa[nums[end]] == 0 {
			unique++
		}
		newmapa[nums[end]]++
		for unique == len(mapa) {
			result += len(nums) - end
			newmapa[nums[start]]--
			if newmapa[nums[start]] == 0 {
				unique--
			}
			start++
		}
	}
	return result
}