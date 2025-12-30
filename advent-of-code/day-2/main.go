package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main_part_1() {

	//Load file in
	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("Failed to open file with error %s", err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		l := s.Text()

		if l == "" {
			log.Fatalf("Blank line")
		}

		t := strings.Split(l, ",")

		for _, s := range t {

			if s == "" {
				continue
			}
			r := strings.Split(s, "-")
			start, err := strconv.Atoi(r[0])
			if err != nil {
				log.Fatalf("Failed to parse range: %c", s[0])
			}
			end, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatalf("Failed to parse range: %c", s[1])
			}

			for i := start; i <= end; i++ {

				snum := strconv.Itoa(i)
				slen := len(snum)

				if slen%2 == 0 {
					smid := slen / 2
					first_half := snum[:smid]
					last_half := snum[smid:]

					if first_half == last_half {
						fmt.Printf("First: %s, last: %s, num: %d\n", first_half, last_half, i)
						sum += i
					}
					// if i%11 == 0 && len(string(i))%2 != 0 {
					// 	fmt.Println(i)
					// }
				}

			}

		}

	}

	fmt.Printf("Sum Total: %d\n", sum)
}

func main_part_2() {

	//Load file in
	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("Failed to open file with error %s", err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		l := s.Text()

		if l == "" {
			log.Fatalf("Blank line")
		}

		t := strings.Split(l, ",")

		for _, s := range t {

			if s == "" {
				continue
			}
			r := strings.Split(s, "-")
			start, err := strconv.Atoi(r[0])
			if err != nil {
				log.Fatalf("Failed to parse range: %c", s[0])
			}
			end, err := strconv.Atoi(r[1])
			if err != nil {
				log.Fatalf("Failed to parse range: %c", s[1])
			}

			for i := start; i <= end; i++ {

				snum := strconv.Itoa(i)
				slen := len(snum)

				// Could optimize for cases like 111 vs 1111 since we cant just do modulo 2
				for j := slen - 1; j > 0; j-- {
					pattern := snum[:j]
					repeated := strings.Repeat(pattern, slen/len(pattern))

					if snum == repeated {
						fmt.Printf("pattern %s, repeated: %s, match: %v\n", pattern, repeated, snum == repeated)
						sum += i
						break
					}
				}

			}

		}

	}

	fmt.Printf("Sum Total: %d\n", sum)
}

func main() {
	main_part_2()
}
