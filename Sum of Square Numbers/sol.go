func judgeSquareSum(c int) bool {
	for a := 1; a < 10; a++ {
		fmt.Println(c - pow(a))
		if c-pow(a) >= 0 && math.Mod(math.Sqrt(float64(c-pow(a))), 1) == 0 {
			return true
		}
	}
	return false
}
func pow(a int) int {
	a *= a
	return a
}
