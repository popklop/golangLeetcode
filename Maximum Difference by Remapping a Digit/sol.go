func minMaxDifference(num int) int {
	mas := []int{}
	for num > 0 {
		mas = append(mas, num%10)
		num = num / 10
	}
	minmas := make([]int, len(mas))
	slices.Reverse(mas)
	copy(minmas, mas)
	m := mas[0]
	n := minmas[0]
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
	for i := 0; i < len(minmas); i++ {
		if minmas[i] != 0 && minmas[i] == n {
			minmas[i] = 0
		}
	}
	fmt.Println(mas, minmas, trans(mas)-trans(minmas))
	return trans(mas)-trans(minmas)
}

func trans(a []int) int {
	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i] * Pow(10, (len(a)-1)-i)
	}

	return result
}
func Pow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y < 0 {
		return 0
	}
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}