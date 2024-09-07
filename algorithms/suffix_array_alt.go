package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// compare the performance between the two
// O(n*n*log n)
func suffix_array_builtin(s string) []int {
	defer timeTrack(time.Now(), "suffix_array_naive")
	type suffix_entry struct {
		suffix string
		pos    int
	}
	strlen := len([]rune(s))
	sx := make([]suffix_entry, strlen)
	result := make([]int, strlen)
	for i := 0; i < strlen; i++ {
		// This UTF-8 syntax is optimized in recent versions of Golang.
		sx[i].suffix = string([]rune(s)[i:])
		sx[i].pos = i
	}
	slices.SortFunc(sx, func(a, b suffix_entry) int {
		return cmp.Compare(a.suffix, b.suffix)
	})
	for i := 0; i < strlen; i++ {
		result[i] = sx[i].pos
	}
	return result
}

// O(n*log n), but more memory efficient and UTF-8 compatible
func suffix_array_linear_two(s string) []int {
	defer timeTrack(time.Now(), "suffix_array_optimized")
	type suffix struct {
		index int
		rank  [2]int
	}
	strlen := len([]rune(s))
	sx := make([]suffix, strlen)
	for i := 0; i < strlen; i++ {
		sx[i].index = i
		sx[i].rank[0] = int([]rune(s)[i]) - int('a')
		if i+1 < strlen {
			sx[i].rank[1] = int([]rune(s)[i+1]) - int('a')
		} else {
			sx[i].rank[1] = -1
		}
	}

	slices.SortFunc(sx, func(a, b suffix) int {
		if a.rank[0] == b.rank[0] {
			return cmp.Compare(a.rank[1], b.rank[1])
		} else {
			return cmp.Compare(a.rank[0], b.rank[0])
		}
	})

	indices := make([]int, strlen)

	for k := 4; k < 2*strlen; k *= 2 {
		rank := 0
		prev_rank := sx[0].rank[0]
		sx[0].rank[0] = rank
		indices[sx[0].index] = 0

		for i := 1; i < strlen; i++ {
			if sx[i].rank[0] == prev_rank && sx[i].rank[1] == sx[i-1].rank[1] {
				prev_rank = sx[i].rank[0]
				sx[i].rank[0] = rank
			} else {
				prev_rank = sx[i].rank[0]
				sx[i].rank[0] = rank + 1
				rank++
			}
			indices[sx[i].index] = i
		}

		for i := 0; i < strlen; i++ {
			nextindex := sx[i].index + (k >> 1)
			if nextindex < strlen {
				sx[i].rank[1] = sx[indices[nextindex]].rank[0]
			}
		}

		slices.SortFunc(sx, func(a, b suffix) int {
			if a.rank[0] == b.rank[0] {
				return cmp.Compare(a.rank[1], b.rank[1])
			} else {
				return cmp.Compare(a.rank[0], b.rank[0])
			}
		})
	}

	suffixarr := make([]int, strlen)

	for i := 0; i < strlen; i++ {
		suffixarr[i] = sx[i].index
	}

	return suffixarr
}

type sarr func(s string) []int

func suffix_array_benchmark(st []string, fn sarr) {
	for _, s := range st {
		//fmt.Printf("processing string №%d...", i+1)
		g := fn(s)
		fmt.Printf("Length is %d", len(g))
		for _, j := range g {
			fmt.Printf("%d; ", j)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// text_file_on_linux: 230 words
	// larger_text_file_on_linux: 2480 words
	// building a performance graph
	for i := 1; i < 8; i *= 2 {
		s := fmt.Sprintf("performance_testing_data/file_%d.txt", i)
		input_file, err := os.ReadFile(s)
		if err != nil {
			fmt.Print("Error reading file, benchmark stopped")
		}
		test_batch := []string{string(input_file)}
		suffix_array_benchmark(test_batch, suffix_array_builtin)
		suffix_array_benchmark(test_batch, suffix_array_linear_two)
	}

	//test_batch := strings.Split(input_string, " ")

	// works even with non-ASCII characters, enabling use of them as isolating characters
	testing_strings := []string{"ababa", "mississippi$", "language", "你好世界", "русский"}
	// works as intended

	suffix_array_benchmark(testing_strings, suffix_array_linear_two)
	// works as intended

}
