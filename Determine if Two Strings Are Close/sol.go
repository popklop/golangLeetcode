func closeStrings(word1 string, word2 string) bool {
	st := make(map[byte]int)
	dn := make(map[byte]int)
	if len(word1) != len(word2) {
		return false
	}
	for i := 0; i < len(word1); i++ {
		st[(word1[i])]++
		dn[(word2[i])]++
	}
	for k, v := range st {
		if v != 0 && dn[k] != 0 {

		} else {
			return false
		}
	}
	temp1 := make(map[int]int)
	temp2 := make(map[int]int)
	for _, v := range dn {
		temp2[v]++
	}
	for _, v := range st {
		temp1[v]++
	}
	for k, v := range temp1 {
		if v2, ok := temp2[k]; !ok || v2 != v {
			return false
		}
	}
	return true
}