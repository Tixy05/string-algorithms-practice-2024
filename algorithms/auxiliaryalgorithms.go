package algorithms

// Memory consumption was sacraficed for simplicity/readability/sanity reasons
// (but still O(n) theoretically :))
func innnerSuffixArray(s []rune, suffixArray []int, alphabetSize int) { //nolint:all
	isSString := make([]bool, len(s)+1)

	isLMSChar := make([]bool, len(s)+1)

	LMSPrefix := func(i int) int {
		if i == len(s) {
			return i
		}
		k := i + 1
		for !isLMSChar[k] {
			k++
		}
		return k
	}

	LMSCompare := func(i, j int) int {
		if i == j {
			return 0
		}
		iEnd := LMSPrefix(i)
		jEnd := LMSPrefix(j)

		for k := 0; k < min(iEnd-i, jEnd-j)+1; k++ {
			if i+k == len(s) {
				return 1
			} else if j+k == len(s) {
				return -1
			}
			if s[i+k] > s[j+k] {
				return 1
			} else if s[i+k] < s[j+k] {
				return -1
			}

			if isSString[i+k] && !isSString[j+k] {
				return 1
			} else if !isSString[i+k] && isSString[j+k] {
				return -1
			}
		}

		if iEnd-i > jEnd-j {
			return 1
		} else if iEnd-i < jEnd-j {
			return -1
		} else {
			return 0
		}
	}

	p1 := []int{}

	charCount := make([]int, alphabetSize)
	for _, c := range s {
		charCount[c]++
	}
	charCount[0] = 1

	var bucketsHeads, bucketsTails []int

	InitBuckets := func() {
		bucketsHeads = make([]int, alphabetSize)
		bucketsTails = make([]int, alphabetSize)
		bucketsHeads[0] = 0
		for i := 1; i < len(bucketsHeads)-1; i++ {
			bucketsHeads[i] = charCount[i-1]
		}
		for i := 2; i < len(bucketsHeads); i++ {
			bucketsHeads[i] += bucketsHeads[i-1]
		}

		for i := 0; i < len(bucketsHeads)-1; i++ {
			bucketsTails[i] = bucketsHeads[i+1] - 1
		}
		bucketsTails[len(bucketsTails)-1] = len(s)
	}

	InitBuckets()

	for i := range suffixArray {
		suffixArray[i] = -1
	}

	isSString[len(s)] = true
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] < s[i+1] || (s[i] == s[i+1] && isSString[i+1]) {
			isSString[i] = true
		}
	}

	for i := len(isSString) - 1; i >= 1; i-- {
		if isSString[i] && !isSString[i-1] {
			isLMSChar[i] = true
			p1 = append(p1, i)
		}
	}

	for i := len(p1) - 1; i >= 0; i-- {
		v := p1[i]
		bucketTailsIndex := 0
		if v != len(s) {
			bucketTailsIndex = int(s[v])
		}
		indexSA := bucketsTails[bucketTailsIndex]
		suffixArray[indexSA] = v
		bucketsTails[bucketTailsIndex]--
	}

	for i, j := 0, len(p1)-1; i < j; i, j = i+1, j-1 {
		p1[i], p1[j] = p1[j], p1[i]
	}

	for _, v := range suffixArray {
		if v <= 0 {
			continue
		}
		if !isSString[v-1] { // if s[v-1] is L-type
			indexSA := bucketsHeads[int(s[v-1])]
			suffixArray[indexSA] = v - 1
			bucketsHeads[int(s[v-1])]++
		}
	}

	InitBuckets()

	for i := len(suffixArray) - 1; i >= 0; i-- {
		v := suffixArray[i]
		if v <= 0 {
			continue
		}
		if isSString[v-1] {
			indexSA := bucketsTails[s[v-1]]
			suffixArray[indexSA] = v - 1
			bucketsTails[s[v-1]]--
		}
	}

	isAllUnique := true
	LMSNames := make(map[int]int32)
	indexName := int32(0)
	for i := 0; i < len(suffixArray)-1; i++ {
		v1 := suffixArray[i]
		v2 := suffixArray[i+1]
		if !isLMSChar[v1] {
			continue
		}
		LMSNames[v1] = indexName
		indexName++
		if isLMSChar[v2] && LMSCompare(v1, v2) == 0 {
			indexName--
			isAllUnique = false
		}
	}
	last := suffixArray[len(suffixArray)-1]
	if isLMSChar[last] {
		LMSNames[last] = indexName
	}
	newAlphabetSize := int(indexName) + 1

	suffixArray1 := make([]int, len(p1))
	s1 := make([]rune, len(p1))
	for i, v := range p1 {
		s1[i] = LMSNames[v]
	}
	s1 = s1[:len(s1)-1]

	if !isAllUnique {
		innnerSuffixArray(s1, suffixArray1, newAlphabetSize)
	} else {
		for i, v := range p1 {
			suffixArray1[LMSNames[v]] = i
		}
	}

	// Induce SA from SA1
	for i := range suffixArray {
		suffixArray[i] = -1
	}

	InitBuckets()

	for i := len(suffixArray1) - 1; i >= 0; i-- {
		v := p1[suffixArray1[i]]
		indexSA := 0
		if v != len(s) {
			indexSA = bucketsTails[s[v]]
		}
		suffixArray[indexSA] = v
		if v != len(s) {
			bucketsTails[s[v]]--
		} else {
			bucketsTails[0]--
		}
	}

	for _, v := range suffixArray {
		if v <= 0 {
			continue
		}
		if !isSString[v-1] { // if s[v-1] is L-type
			indexSA := bucketsHeads[int(s[v-1])]
			suffixArray[indexSA] = v - 1
			bucketsHeads[int(s[v-1])]++
		}
	}

	InitBuckets()

	for i := len(suffixArray) - 1; i >= 0; i-- {
		v := suffixArray[i]
		if v <= 0 {
			continue
		}
		if isSString[v-1] {
			indexSA := bucketsTails[s[v-1]]
			suffixArray[indexSA] = v - 1
			bucketsTails[s[v-1]]--
		}
	}
}

func SuffixArray(s []rune, alphabetSize int) []int {
	if len(s) == 0 {
		return []int{}
	}
	suffixArray := make([]int, len(s)+1)
	innnerSuffixArray(s, suffixArray, alphabetSize)
	return suffixArray[1:] // first element of SA is always sentinel
}

// TODO: make an independent alphabetSize() function and give user
// an option to select to compute this in main SuffixArray function

// Algorithm was taken from
// "Two Efﬁcient Algorithms for Linear Time Sufﬁx Array Construction"
// by Sen Zhang and Daricks Wai Hong Chan
// First algorithm was chosen (SA-IS)
// www.researchgate.net/publication/224176324
