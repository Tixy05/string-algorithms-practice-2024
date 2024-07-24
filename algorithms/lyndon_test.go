package algorithms

import (
	"reflect"
	"testing"
)

func TestLyndonFactorization(t *testing.T) {
	factorizations := map[string]([]string){
		"abcdef":         []string{"abcdef"},
		"bababa":         []string{"b", "ab", "ab", "a"},
		"aaa":            []string{"a", "a", "a"},
		"abbaabbbaaabab": []string{"abb", "aabbb", "aaabab"},
	}

	for str, factorization := range factorizations {
		result := LyndonFactorization(str)
		if !reflect.DeepEqual(result, factorization) {
			t.Fatalf("wrong Lyndon factorization of \"%s\"\n\twant: %v\n\tgot: %v", str, factorization, result)
		}
	}
}
