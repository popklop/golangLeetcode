func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result []int
	result = append(result, nums1...)
	result = append(result, nums2...)
	var s float64
	mas := sort(result)
	if len(mas)%2 != 0 {
		s = float64(mas[len(result)/2.0])
	} else {
		s = float64(mas[len(result)/2.0]+mas[len(result)/2-1]) / 2
	}
	return s
}

func sort(arr []int) []int {
	mid := len(arr) / 2
	if len(arr) <= 1 {
		return arr
	}
	masl := sort(arr[:mid])
	masr := sort(arr[mid:])
	return merge(masl, masr)
}

func merge(nums1 []int, nums2 []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	result = append(result, nums1[i:]...)
	result = append(result, nums2[j:]...)
	return result
}
