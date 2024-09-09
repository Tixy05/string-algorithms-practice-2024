package algorithms

type lyndonSubstring struct {
	start, end int
}

func LyndonFactorization(s []rune) []lyndonSubstring {
	res := []lyndonSubstring{}
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
			res = append(res, lyndonSubstring{i, i + k - j - 1})
			i += k - j
		}
	}

	return res
}
