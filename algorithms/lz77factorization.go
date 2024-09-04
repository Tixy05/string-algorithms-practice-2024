package algorithms

type lz77factor struct {
	startOrSymbol int32
	length        int32
}

func LZ77Factorization(s []rune) []lz77factor {
	LCP := func(i, j int) int { // Longest Common Prefix
		if i == 0 || j == 0 {
			return 0
		}
		i--
		j--
		var k int
		for k = 0; max(i+k, j+k) < len(s) && s[i+k] == s[j+k]; k++ {
		}
		return k
	}

	LZ77Factor := func(i, psv, nsv int) (int, lz77factor) {
		lcp1 := LCP(i, psv)
		lcp2 := LCP(i, nsv)
		var res lz77factor
		if lcp1 > lcp2 {
			res = lz77factor{int32(psv), int32(lcp1)}
		} else {
			res = lz77factor{int32(nsv), int32(lcp2)}
		}
		if res.length == 0 {
			res.startOrSymbol = s[i-1]
		} else {
			res.startOrSymbol--
		}
		return i + int(max(res.length, 1)), res
	}
	maxC := rune(0)
	for _, c := range s {
		if c > maxC {
			maxC = c
		}
	}
	maxC++
	tmpSA := SuffixArray(s, int(maxC)+1)
	for i := range tmpSA {
		tmpSA[i]++
	}

	SA := []int{0}
	SA = append(SA, tmpSA...)
	SA = append(SA, 0)
	NSV := make([]int, len(SA)) // Next Smaller Values
	PSV := make([]int, len(SA)) // Previous Smaller Values
	top := 0
	for i := 1; i <= len(s)+1; i++ {
		for SA[top] > SA[i] {
			NSV[SA[top]] = SA[i]
			PSV[SA[top]] = SA[top-1]
			top--
		}
		top++
		SA[top] = SA[i]
	}

	res := []lz77factor{}
	var f lz77factor
	i := 1
	for i <= len(s) {
		i, f = LZ77Factor(i, PSV[i], NSV[i])
		res = append(res, f)
	}
	return res
}

func FromLZ77FactorizationToString(factorization []lz77factor) string {
	length := 0
	for _, v := range factorization {
		if v.length == 0 {
			length++
		} else {
			length += int(v.length)
		}
	}
	res := make([]rune, length)
	strIndex := 0
	for _, v := range factorization {
		if v.length == 0 {
			res[strIndex] = v.startOrSymbol
			strIndex++
		} else {
			start := v.startOrSymbol
			for i := 0; i < int(v.length); i++ {
				res[strIndex] = res[start+int32(i)]
				strIndex++
			}
		}
	}
	return string(res)
}

// Algorithm was taken from
// "Linear Time Lempel-Ziv Factorization: Simple, Fast, Small"
// by Juha Kärkkäinen, Dominik Kempa, and Simon J. Puglisi
