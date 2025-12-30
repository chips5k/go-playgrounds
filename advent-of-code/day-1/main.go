package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := 50
	ppos := start
	pos := start
	max := 100
	dir := ""
	count := 0

	tot_rots := 0
	for scanner.Scan() {
		ln := scanner.Text()

		if ln[0] != 'L' && ln[0] != 'R' {
			log.Fatalf("Invalid direction: %s", ln)
		}

		rots := 0
		amt, err := strconv.Atoi(ln[1:])

		if err != nil {
			log.Fatalf("Invalid amount: %s", ln)
		}

		switch ln[0] {
		case 'R':
			dir = "Right"
			pos += amt
			rots = pos / max

		case 'L':
			dir = "Left"
			pos -= amt

			if ppos == 0 {
				rots = amt / max
			} else if amt >= ppos {
				rots = 1 + (amt-ppos)/max
			}

		}

		pos = (pos%max + max) % max

		if pos == 0 {
			count++
		}

		fmt.Printf("Move %s by %d from %d to %d, rots: %d\n", dir, amt, ppos, pos, rots)
		ppos = pos
		tot_rots += rots
	}

	fmt.Printf("Part 1 Code is %d, Part 2 Code is %d, should be 992 & 6133\n", count, tot_rots)

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}
