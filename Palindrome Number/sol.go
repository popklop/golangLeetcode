func isPalindrome(x int) bool {
     

	var mas []int
    if x<0{
        return false
    }
	for x > 0 {
		mas = append(mas, x%10)
		x = x / 10
	}
	revmas := revarr(mas)
	for i := range mas {
		if mas[i] != revmas[i] {
			return false
			break
		}

	}
	return true
}


func revarr(mas []int) []int {
	var revedmas []int
	//var i int =
	for i := len(mas) - 1; i >= 0; i-- {
		revedmas = append(revedmas, (mas)[i])
	}
	return revedmas
}