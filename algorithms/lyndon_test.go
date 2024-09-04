package algorithms

import (
	"reflect"
	"testing"
)

func TestLyndonFactorization(t *testing.T) {
	factorizations := map[string]([]lyndonSubstring){
		"abcdef":         []lyndonSubstring{{0, 5}},                         // abcdef
		"bababa":         []lyndonSubstring{{0, 0}, {1, 2}, {3, 4}, {5, 5}}, // b ab ab a
		"aaa":            []lyndonSubstring{{0, 0}, {1, 1}, {2, 2}},         // a a a
		"abbaabbbaaabab": []lyndonSubstring{{0, 2}, {3, 7}, {8, 13}},        // abb aabbb aaabab
	}

	for str, factorization := range factorizations {
		result := LyndonFactorization([]rune(str))
		if !reflect.DeepEqual(result, factorization) {
			t.Fatalf("wrong Lyndon factorization of \"%s\"\n\twant: %v\n\tgot: %v", str, factorization, result)
		}
	}
}
