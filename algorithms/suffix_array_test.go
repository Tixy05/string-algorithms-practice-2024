package algorithms

import (
	"reflect"
	"testing"
)

func TestAlternativeSuffixArray(t *testing.T) {
	TestData := map[string]([]int) {
		"mississippi":  []int{10, 7, 4, 1, 0, 9, 8, 6, 3, 5, 2},
		"wolloomooloo": []int{2, 9, 3, 6, 11, 1, 8, 5, 10, 7, 4, 0},
		"abcdef":       []int{0, 1, 2, 3, 4, 5},
		"fedcba":       []int{5, 4, 3, 2, 1, 0},
		"":             []int{},
		"русский":	[]int{5, 6, 4, 0, 3, 2, 1}
	}

	for str, sa := range TestData {
		// naive suffix array variant
		result := suffix_array_builtin(str)
		if !reflect.DeepEqual(result, sa) {
			t.Fatalf("wrong suffix array of \"%s\"\n\twant: %v\n\tgot: %v", str, sa, result)
		}
		// optimized alternative suffix array
		// Pros:
		// UTF-8 compatible
		// no need to specify alphabet size
		// Cons:
		// theroretical performance O(n*log n)
		result := suffix_array_linear_two(str)
		if !reflect.DeepEqual(result, sa) {
			t.Fatalf("wrong suffix array of \"%s\"\n\twant: %v\n\tgot: %v", str, sa, result)
		}
	}
}
