package algorithms

import (
	"reflect"
	"testing"
)

func TestCrochemore(t *testing.T) {
	TestData := map[string]([]string){
		"ababa":        []string{"abab"},
		"acababaeeaee": []string{"ee", "ee", "abab", "baba", "aeeaee"},
		"dfssdfdf":     []string{"ss", "dfdf"},
		"яруярусский":  []string{"сс", "яруяру"},
		"我天天喝茶":        []string{"天天"},
	}

	for str, sa := range TestData {

		result := crochemore_decomposition_simplified(str)
		if !reflect.DeepEqual(result, sa) {

			for i, a := range result {
				if a == sa[i] {
					t.Fatalf("wrong crochemore decomposition of \"%s\"\n\twant: %v\n\tgot: %v", str, sa[i], a)
				}
			}

		}

	}
}
