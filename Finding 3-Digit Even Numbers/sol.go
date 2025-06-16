func findEvenNumbers(digits []int) []int {
	arr := []int{}
	for i := 0; i < len(digits); i++ {
		if digits[i]%2 == 0 {
			arr = append(arr, i)
		}
	}
	mapa := make(map[int]int)
	result := []int{}
	for h := 0; h < len(arr); h++ {
		for i := 0; i < len(digits); i++ {
			for j := 0; j < len(digits); j++ {
				if i != j && i != arr[h] && j != arr[h] && digits[i] != 0 {
					threeDigitNum := digits[i]*100 + digits[j]*10 + digits[arr[h]]
					if mapa[threeDigitNum] == 0 {
						mapa[threeDigitNum]++
						result = append(result, threeDigitNum)
					}

				}
			}
		}
	}
slices.Sort(result)
	return result
}