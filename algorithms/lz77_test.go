package algorithms

import (
	"testing"
)

func TestLZ77(t *testing.T) {
	strs := []string{
		"abcdef",
		"bababa",
		"aaa",
		"abbaabbbaaabab",
	}

	for _, str := range strs {
		result := FromLZ77FactorizationToString(
			LZ77Factorization([]rune(str)),
		)
		if result != str {
			t.Fatalf("string induced from factorization do not"+
				"equals to original string"+
				"\n\twant: %v\n\tgot: %v", str, result)
		}
	}
}
