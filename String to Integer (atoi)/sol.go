func myAtoi(s string) int {
	var alph []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", " ", "+"}
	var testmas []int
	tobemin := false
	res := 0
	was := false

	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case alph[0]:
			testmas = append(testmas, 0)
		case alph[1]:
			testmas = append(testmas, 1)
		case alph[2]:
			testmas = append(testmas, 2)
		case alph[3]:
			testmas = append(testmas, 3)
		case alph[4]:
			testmas = append(testmas, 4)
		case alph[5]:
			testmas = append(testmas, 5)
		case alph[6]:
			testmas = append(testmas, 6)
		case alph[7]:
			testmas = append(testmas, 7)
		case alph[8]:
			testmas = append(testmas, 8)
		case alph[9]:
			testmas = append(testmas, 9)
		case alph[10]: 
			if len(testmas) == 0 && !was {
				tobemin = true
				was = true
			} else {
				i = len(s)
			}
		case alph[11]: 
			if len(testmas) != 0 || was {
				i = len(s)
			}
			break
		case alph[12]: 
			if len(testmas) == 0 && !was {
				was = true
			} else {
				i = len(s)
			}
		default:
	
			i = len(s)
		}
	}

	
	if len(testmas) == 0 {
		return 0
	}

	
	for i := 0; i < len(testmas); i++ {
		if res > (2147483647-testmas[i])/10 {
			if tobemin {
				return -2147483648
			}
			return 2147483647
		}
		res = res*10 + testmas[i]
	}

	
	if tobemin {
		res = -res
	}

	
	if res < -2147483648 {
		return -2147483648
	} else if res > 2147483647 {
		return 2147483647
	}

	return res
}