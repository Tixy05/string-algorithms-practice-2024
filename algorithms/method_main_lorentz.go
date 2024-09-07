package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
	"log"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func z_array(s string) []int {
	n := len([]rune(s))
	z := make([]int, n)
	runes := []rune(s)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && runes[z[i]] == runes[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}
	// return the reversed string.
	return string(rns)
}

func print_tandem(s string, shift int, cond bool, cntr int, l int, l1 int, l2 int) {
	var pos int
	if cond {
		pos = cntr - l1
	} else {
		pos = cntr - 2*l1 - l2 + 1
	}
	op_p := string("[")
	im := string(";")
	cl_p := string("]")
	space := string(" ")
	eq := string("=")
	left := strconv.Itoa(shift + pos)
	right := strconv.Itoa(shift + pos + 2*l - 1)
	//sstr := s[pos : pos+2*l]
	sstr := string([]rune(s)[pos : pos+2*l])
	prompt := op_p + left + im + right + cl_p + space + eq + space + sstr
	fmt.Println(prompt)
}

func print_tandems(s string, shift int, cond bool, cntr int, l int, k1 int, k2 int) {
	for l1 := 1; l1 <= l; l1++ {
		if cond && l1 == l {
			break
		}
		if l1 <= k1 && l-l1 <= k2 {
			print_tandem(s, shift, cond, cntr, l, l1, l-l1)
		}
	}
}

func check_z(z []int, i int) int {
	if i < 0 || i >= len(z) {
		return 0
	}
	return z[i]
}

func main_lorentz_tandems(s string, shift int) {

	strlen := len([]rune(s))
	if strlen == 1 {
		return
	}
	half := strlen / 2
	rhalf := strlen - half
	lh := string([]rune(s)[0:half])
	rh := string([]rune(s)[half:])
	lh_r := reverse(lh)
	rh_r := reverse(rh)
	main_lorentz_tandems(lh, shift)
	main_lorentz_tandems(rh, shift+half)

	isolation_char := string("𡾌")
	z1_array := z_array(lh_r)
	z2_array := z_array(rh + isolation_char + lh)
	z3_array := z_array(lh_r + isolation_char + rh_r)
	z4_array := z_array(rh)

	for cntr := 0; cntr < strlen; cntr++ {
		var l int
		var k1 int
		var k2 int

		if cntr < half {
			l = half - cntr
			k1 = check_z(z1_array, half-cntr)
			k2 = check_z(z2_array, rhalf+cntr+1)
		} else {
			l = cntr - half + 1
			k1 = check_z(z3_array, 2*half+rhalf-cntr)
			k2 = check_z(z4_array, cntr-half+1)
		}

		if k1+k2 >= 1 {
			l = l+1
			return
			//print_tandems(s, shift, cntr < half, cntr, l, k1, k2)
		}

	}

}

func tandem_benchmark(mass []string) {
	defer timeTrack(time.Now(), "main_lorentz")
	for _, s := range mass {
		//fmt.Printf("Word №%d:\n", i)
		main_lorentz_tandems(s, 0)
	}
}


// word length := [lmin; lmax)
func rand_word_generator(lmin int, lmax int, words int) []string {
	var result []string
	for i := 0; i < words; i++ {

		wordlen := lmin + rand.Intn(lmax-lmin)
		var curword string
		for j := 0; j < wordlen; j++ {
			// 65 == 'A'
			// 122 == 'z'
			nextchar := 65 + rand.Intn(122-65+1)
			curword = curword + string(nextchar)
		}
		result = append(result, curword)

	return result
}

func main() {
	// bm := []string{"ababa", "acababaeeaee", "dfssdfdf", "яруярусский", "我天天喝茶"}

	fmt.Printf("Choose benchmark type: \n(1) Real-world text benchmark;\n(2) Random word benchmark: ")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		log.Printf("Wrong choice number!")
		return
	}
	if choice == 1 {
		for i := 1; i < 8; i *= 2 {
			s := fmt.Sprintf("performance_testing_data/file_%d.txt", i)
			input_file, err := os.ReadFile(s)
			if err != nil {
				fmt.Print("Error reading file, benchmark stopped")
			}
			test_batch := strings.Split(string(input_file), " ")
			tandem_benchmark(test_batch)
		}
	} else if choice == 2 {
		var minl int
		fmt.Printf("Enter min word length: ")
		_, err = fmt.Scanf("%d", &minl)
		if err != nil {
			log.Printf("Error word length!")
			return
		}
		if minl <= 0 {
			log.Printf("Non-positive value for word length!")
			return
		}

		fmt.Printf("Enter max word length: ")
		var maxl int
		_, err = fmt.Scanf("%d", &maxl)
		if err != nil {
			log.Printf("Error word length!")
			return
		}
		if maxl <= 0 {
			log.Printf("Non-positive value for word length!")
			return
		}
		if maxl <= minl {
			log.Printf("Value for maximum word length can't be less than min word length!")
			return
		}

		fmt.Printf("Enter number of words: ")
		var words int
		_, err = fmt.Scanf("%d", &words)
		if err != nil {
			log.Printf("Error word length!")
			return
		}
		if words <= 0 {
			log.Printf("Non-positive value for word count!")
			return
		}

		tandem_benchmark(rand_word_generator(minl,maxl,words))

	}



	//var reps int
	//var bm []string
	//fmt.Printf("Enter number of test strings: ")
	//_, err := fmt.Scanf("%d", &reps)
	//if err != nil || reps <= 0 {
	//	fmt.Println("Number is not valid")
	//	return
	//}

	//for i := range reps {
	//	reader := bufio.NewReader(os.Stdin)
	//	fmt.Printf("[%d] Enter text: ", i+1)
	//	text, _ := reader.ReadString('\n')
	//	bm = append(bm, text)
	//}
	//tandem_benchmark(bm)

}
