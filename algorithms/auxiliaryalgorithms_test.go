package algorithms

import (
	"reflect"
	"testing"
)

func TestSA(t *testing.T) {
	suffixArrays := map[string]([]int){
		"mississippi":  []int{10, 7, 4, 1, 0, 9, 8, 6, 3, 5, 2},
		"wolloomooloo": []int{2, 9, 3, 6, 11, 1, 8, 5, 10, 7, 4, 0},
		"abcdef":       []int{0, 1, 2, 3, 4, 5},
		"fedcba":       []int{5, 4, 3, 2, 1, 0},
		"":             []int{},
	}
	s := "the quick brown fox jumps over the lazy dog"
	expected := []int{9, 39, 15, 19, 34, 25, 3, 30, 36, 10, 7, 40, 33,
		2, 28, 16, 42, 32, 1, 6, 20, 8, 35, 22, 14, 41, 26, 12, 17,
		23, 4, 29, 11, 24, 31, 0, 5, 21, 27, 13, 18, 38, 37}
	suffixArrays[s] = expected

	for str, sa := range suffixArrays {
		result := SuffixArray([]rune(str), 10000)
		if !reflect.DeepEqual(result, sa) {
			t.Fatalf("wrong suffix array of \"%s\"\n\twant: %v\n\tgot: %v", str, sa, result)
		}
	}
}

// Special thanks to Lloyd Allison and his
// website www.allisons.org/ll/ for allowing to compute
// SuffixArray online and steal result fro testing purposes
