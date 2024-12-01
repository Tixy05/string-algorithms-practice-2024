package algorithms

import (
	"reflect"
	"testing"
)

func TestSuffArray(t *testing.T) {
	TestData := map[string]([]int){
		"ababa":        []int{4, 2, 0, 3, 1},
		"acababaeeaee": []int{2, 4, 0, 9, 6, 3, 5, 1, 11, 8, 10, 7},
		"dfssdfdf":     []int{6, 4, 0, 7, 5, 1, 3, 2},
		"яруярусский":  []int{9, 10, 8, 4, 1, 7, 6, 5, 2, 3, 0},
		"我天天喝茶":        []int{3, 2, 1, 0, 4},
	}

	for str, sa := range TestData {
		// naive suffix array variant
		result, time := suffix_array_builtin(str)
		if !reflect.DeepEqual(result, sa) {
			t.Fatalf("wrong suffix array of \"%s\"\n\twant: %v\n\tgot: %v", str, sa, result)
		} else {
			t.Logf("Test on %s executed successfully in %d μs", str, time)
		}
		// optimized alternative suffix array
		// Pros:
		// practically unlimited alphabet size
		// no need to specify alphabet
		// Cons:
		// theroretical performance O(n*log n)
		result, time = suffix_array_linear_two(str)
		if !reflect.DeepEqual(result, sa) {
			t.Fatalf("wrong suffix array of \"%s\"\n\twant: %v\n\tgot: %v", str, sa, result)
		} else {
			t.Logf("Test on %s executed successfully in %d μs", str, time)
		}
	}
}
