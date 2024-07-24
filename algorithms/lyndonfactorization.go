package algorithms

func LyndonFactorization(s string) []string {
	res := []string{}
	n := len(s)
	i := 0
	for i < n {
		j := i
		k := i + 1
		for k < n && s[j] <= s[k] {
			if s[j] < s[k] {
				j = i
			} else {
				j++
			}
			k++
		}
		for i <= j {
			res = append(res, s[i:i+k-j])
			i += k - j
		}
	}

	return res
}
