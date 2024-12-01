package algorithms

// EXTREMELY simplified decomposiition algorithm for searching tandem string repetitions
// On the plus side, works with any alphabet
func crochemore_decomposition_simplified(str string) []string {
	runes := []rune(str)
	str_len := len(runes)
	var result []string
	for level := 1; level <= str_len/2; level++ {
		map3 := make(map[string][]int)
		for j := 0; j <= str_len-level; j++ {
			test_substr := runes[j : j+level]
			map3[string(test_substr)] = append(map3[string(test_substr)], j)
			//fmt.Printf("%s\n", string(test_substr))
		}
		for str, val := range map3 {
			for i, v := range val {
				for j := i + 1; j < len(val); j++ {
					if val[j]-v == level {

						result = append(result, str+str)
						//fmt.Printf("[%d,%d] = %s%s\n", v, val[j], str, str)
						//if strmax_len != 2*level {
						//	strmax = str + str
						//}
						//strmax_len = 2 * level
						break
					}
				}
			}
		}
	}

	return result
}
