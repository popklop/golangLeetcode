func minSum(nums1 []int, nums2 []int) int64 {
	nulls1 := 0
	nulls2 := 0
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == 0 {
			nulls1++
		}

	}
	for i := 0; i < len(nums2); i++ {
		if nums2[i] == 0 {
			nulls2++
		}
	}
	sums1 := zapolnnull(nums1)
	sums2 := zapolnnull(nums2)
	if (sums1 > sums2) && nulls2 == 0 {
		return -1
	}
	if (sums2 > sums1) && nulls1 == 0 {
		return -1
	}
	if sums1 == sums2 {
		return sums1
	}
	if sums2 > sums1 {
		return sums2
	}
	if sums1 > sums2 {
		return sums1
	}

	return 1
}

func zapolnnull(arr []int) int64 {
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[i] = 1
		}
		sum += arr[i]
	}
	return int64(sum)
}