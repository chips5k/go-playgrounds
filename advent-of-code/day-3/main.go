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

		bb := strings.Split(l, "")

		max := 0
		for i, v := range bb {
			for j := i + 1; j < len(bb); j++ {
				val, err := strconv.Atoi(v + bb[j])
				if err != nil {
					log.Fatalf("Failed to parse %s", v+bb[j])
				}

				if val > max {
					max = val
				}
			}
		}

		sum += max

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

		bb := strings.Split(l, "")

		joltage := ""
		maxSize := 12
		bbSize := len(bb)
		start := 0
		for position := 1; position <= maxSize; position++ {
			curr := 0
			//fmt.Printf("Pos: %d\n==============\n", position)
			remaining := maxSize - position
			//fmt.Printf("Start: %d\n", start)
			for i := start; i < bbSize; i++ {
				v, err := strconv.Atoi(bb[i])
				if err != nil {
					log.Fatalf("Failed to parse value: '%s'", bb[i])
				}
				//fmt.Printf("Is %d > %d and is rb: %d < %d\n", v, curr, remaining, bbSize-i+1)
				if v > curr && remaining < bbSize-i {
					curr = v
					start = i + 1
				}

			}
			//fmt.Printf("Found: %d at %d\n\n", curr, start-1)
			joltage = joltage + strconv.Itoa(curr)
		}

		iJoltage, err := strconv.Atoi(joltage)

		if err != nil {
			log.Fatalf("Failed to parse joltage: '%s'", joltage)
		}

		sum += iJoltage

		fmt.Println(joltage)
		//largest number with at least 11 digits after
		//largest number with at least 10 digits after
		//largest number with at least 9 digits after

	}

	fmt.Printf("Sum Total: %d\n", sum)
}

func main() {
	main_part_2()
}
