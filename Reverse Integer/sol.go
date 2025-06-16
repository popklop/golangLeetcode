func reverse(x int) int {
	var mas []int

	for x > 0 {
		mas = append(mas, x%10)
		x = x / 10
	}
	toret := 0
	for i, num := range mas {
		toret += num * int(math.Pow(10, float64(len(mas)-i-1)))
	}
	if x < 0 {
		x = x * (-1)
		toret = reverse(x) * (-1)
	}
	if toret >= int(math.Pow(-2, 31)) && toret >= int(math.Pow(2, 31))-1 {
		return 0
	}
	return toret
}