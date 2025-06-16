func divide(dividend int, divisor int) int {
	sign := 1

	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}
	dividend = abs(dividend)
	divisor = abs(divisor)
	result := 0
	for dividend >= divisor {
		shift := 0
		for (divisor << shift) <= dividend {
			shift++
		}
		shift--
		result += 1 << shift
		dividend -= divisor << shift
	}
    if result*sign>=2147483648{
        return (result-1)*sign
    }
	return sign * result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}