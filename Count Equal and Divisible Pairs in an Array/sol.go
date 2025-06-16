func countPairs(nums []int, k int) int {
	count := 0
	mapa := make(map[int][]int)

	for i := 0; i < len(nums); i++ {

		for _, j := range mapa[nums[i]] {
			if (i*j)%k == 0 {
				count++
			}
		}
		mapa[nums[i]] = append(mapa[nums[i]], i)
	}

	return count
}