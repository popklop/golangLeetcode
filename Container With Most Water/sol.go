func maxArea(height []int) int {
	res := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		if res < area {
			res = area
		}
	}

	return res
}
