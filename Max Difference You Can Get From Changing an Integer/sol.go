func maxDiff(num int) int {
	mas := []int{}
	for num > 0 {
		mas = append(mas, num%10)
		num = num / 10
	}
	slices.Reverse(mas)
	minmas := make([]int, len(mas))
	copy(minmas, mas)
	m := mas[0]
	was := false
	for i := 0; i < len(mas); i++ {
		if mas[i] == 9 && was == false {
			if i == len(mas)-1 {
				break
			}
			m = mas[i+1]
			continue

		}
		if mas[i] != 9 && mas[i] == m {
			was = true
			mas[i] = 9
		}

	}
	if minmas[0] != 1 {
		firstDigit := minmas[0]
		for i := 0; i < len(minmas); i++ {
			if minmas[i] == firstDigit {
				minmas[i] = 1
			}
		}
	} else {
		minReplace := -1
		for i := 1; i < len(minmas); i++ {
			if minmas[i] != 0 && minmas[i] != 1 {
				minReplace = minmas[i]
				break
			}
		}
		if minReplace != -1 {
			for i := 0; i < len(minmas); i++ {
				if minmas[i] == minReplace {
					minmas[i] = 0
				}
			}
		}
	}

	return trans(mas) - trans(minmas)
}

func trans(a []int) int {
	result := 0
	power := 1
	for i := len(a) - 1; i >= 0; i-- {
		result += a[i] * power
		power *= 10
	}
	return result
}