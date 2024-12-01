package algorithms

import (
	"cmp"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"
)

// compare the performance between the two
// O(n*n*log n)
func suffix_array_builtin(s string) ([]int, int64) {
	t1 := time.Now()
	//defer timeTrack(time.Now(), "suffix_array_naive")
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
	elapsed := time.Since(t1)
	return result, elapsed.Microseconds()
}

// O(n*log n), but more memory efficient and UTF-8 compatible
func suffix_array_linear_two(s string) ([]int, int64) {
	t1 := time.Now()
	//defer timeTrack(time.Now(), "suffix_array_optimized")
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
	elapsed := time.Since(t1)

	return suffixarr, elapsed.Microseconds()
}

type sarr func(s string) ([]int, int64)

func suffix_array_benchmark(st []string, fn sarr) []string {
	time_res := make([]string, len(st))
	for i, s := range st {
		//fmt.Printf("processing string №%d...", i+1)
		_, time := fn(s)
		time_res[i] = fmt.Sprintf("%d", time)
		//fmt.Printf("Length is %d", len(g))
		//for _, j := range g {
		//	fmt.Printf("%d; ", j)
		//}
		//fmt.Printf("\n")
	}
	return time_res
}

func sarr_singular_benchmark(st string, fn sarr, tries int) string {
	//fmt.Printf("processing string №%d...", i+1)
	var time_accum int64
	time_accum = 0
	for i := 0; i < tries; i++ {
		_, time := fn(st)
		time_accum += time
	}
	time_res := time_accum / int64(tries)
	length_s := len(st)
	fmt.Printf("Length is %d, time is %d μs", length_s, time_res)

	//for _, j := range g {
	//	fmt.Printf("%d; ", j)
	//}
	fmt.Printf("\n")
	return fmt.Sprintf("%d", time_res)
}

func urandom_string(length int) string {
	res := make([]byte, length)
	for i := range res {
		// the fastest way to get random characters in ASCII
		res[i] = byte(rand.Int63() % int64(128))
	}
	return string(res)
}

func test_to_file() {
	const gradation = 100
	const times = 5

	f, err := os.Create("performance_graph.csv")

	if err != nil {
		fmt.Print("Error creating file, benchmark stopped")
		return
	}

	r := csv.NewWriter(f)

	header := []string{"length (bytes)", "suffix_array_naive", "suffix_array_optimized"}

	err = r.Write(header)
	if err != nil {
		fmt.Print("Failed to write header to output file")
	}

	for i := 1; i <= 100; i++ {
		testing_string := urandom_string(i * gradation)

		s1 := sarr_singular_benchmark(testing_string, suffix_array_builtin, times)
		s2 := sarr_singular_benchmark(testing_string, suffix_array_linear_two, times)
		err = r.Write([]string{fmt.Sprintf("%d", i*gradation), s1, s2})
		if err != nil {
			fmt.Print("Failed to write value row")
		}
	}
	r.Flush()
	f.Close()
}
