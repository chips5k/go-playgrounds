package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main_part_1() {

	f, err := os.Open("input")

	if err != nil {
		log.Fatalln("Failed to open file")
	}

	scanner := bufio.NewScanner(f)

	var probs [][]string = nil

	for scanner.Scan() {
		l := scanner.Text()

		cols := strings.Fields(l)
		if probs == nil {
			probs = make([][]string, len(cols))
		}
		for i, v := range cols {
			if probs[i] == nil {
				probs[i] = make([]string, 0, 20)
			}
			probs[i] = append(probs[i], v)
		}
	}

	total := 0
	for _, v := range probs {
		amt := 0
		largs := len(v) // ignore the operator
		op := v[largs-1]
		if op == "*" {
			amt = 1
		}

		for j := 0; j < largs-1; j++ {
			arg, err := strconv.Atoi(v[j])

			if err != nil {
				log.Fatalf("Failed to parse: %s", v[j])
			}

			if op == "*" {
				fmt.Printf("%d * %d = %d\n", amt, arg, amt*arg)
				amt *= arg
			} else {
				fmt.Printf("%d + %d = %d\n", amt, arg, amt+arg)
				amt += arg
			}

		}
		total += amt
	}
	fmt.Println(total)
}

func main_part_2_attempt_1() {

	f, err := os.Open("sample")

	if err != nil {
		log.Fatalln("Failed to open file")
	}

	scanner := bufio.NewScanner(f)

	var probs [][]string = nil
	var probWidths []int = nil
	width := 0
	for scanner.Scan() {
		l := scanner.Text()

		cols := strings.Split(l, "")
		fields := strings.Fields(l)

		if probs == nil {
			probs = make([][]string, len(fields))
			probWidths = make([]int, len(fields))
		}

		col := 0
		for i := 0; i < len(cols); i += width + 1 {

			//determine width
			width, err = getWidth(cols[i:])
			fmt.Printf("i = %d, s: %s, Width: %d", i, cols[i:width], width)
			probWidths[col] = width
			if err != nil {
				log.Fatalf("Failed to determine width")
			}

			if probs[col] == nil {
				probs[col] = make([]string, 0, 20)
			}

			joined := strings.Join(cols[i:i+width], "")

			fmt.Printf("'%s'\n", joined)
			//fmt.Println(joined, i)
			probs[col] = append(probs[col], joined)
			col++
		}

	}

	// total := 0

	// for i, prb := range probs {

	// 	//op := strings.Trim(prb[len(prb)-1], " ")

	// 	//dd := make([]string, 10)

	// 	for j := range prb {
	// 		fmt.Printf("prb[%d] = '%s'\n", j, prb[j])
	// 		curr := ""
	// 		for k := 0; k < probWidths[i]; k++ {
	// 			fmt.Printf("prb[%d][%d] = '%c'\n", j, k, prb[j][k])
	// 			curr += string(prb[j][k])
	// 			// 	//fmt.Printf("tok: '%s' len: %d, ind: %d \n", prb[j], len(prb[j]), k)
	// 			// 	dd[k] += string(prb[j][k])
	// 		}
	// 		fmt.Println(curr)

	// 	}

	// 	//fmt.Println(dd)

	// 	//fmt.Println(dd)

	// 	// for i, s := range prb {
	// 	// 	if i < len(prb)-1 {
	// 	// 		for j, c := range s {
	// 	// 			dd[j] += string(c)
	// 	// 		}
	// 	// 	}
	// 	// }

	// 	amt := 0

	// 	// for i, v := range dd {
	// 	// 	if v != "" {
	// 	// 		n, err := strconv.Atoi(v)
	// 	// 		if err != nil {
	// 	// 			log.Fatalf("Failed to parse: %d, '%s', '%s', '%s'", i, v, dd, prb)
	// 	// 		}

	// 	// 		if op == "*" {
	// 	// 			if i == 0 && amt == 0 {
	// 	// 				amt = 1
	// 	// 			}
	// 	// 			amt *= n
	// 	// 			if i != 0 {
	// 	// 				fmt.Printf("*")
	// 	// 			}
	// 	// 			fmt.Printf("%d", n)
	// 	// 		} else {
	// 	// 			if i != 0 {
	// 	// 				fmt.Printf("+")
	// 	// 			}
	// 	// 			amt += n
	// 	// 			fmt.Printf("%d", n)
	// 	// 		}
	// 	// 	}

	// 	// }
	// 	// fmt.Printf(" = %d", amt)
	// 	// fmt.Println()
	// 	total += amt
	// }

	// fmt.Printf("Total: %d\n", total)

}
func getWidth(s []string) (int, error) {

	startChars := false
	for i := 0; i < len(s); i++ {

		if s[i] != " " {
			startChars = true
		}

		if startChars {
			if i+1 >= len(s) {
				return i + 1, nil
			}

			if s[i] == " " && s[i+1] != " " {
				return i, nil
			}

		}
	}

	return -1, errors.New("Failed to determine length")
}
func main_part_2_attempt_2() {

	f, err := os.Open("input")

	if err != nil {
		log.Fatalln("Failed to open input file")
	}

	scanner := bufio.NewScanner(f)
	pstrings := make([]string, 0, 100)

	for scanner.Scan() {
		pstrings = append(pstrings, scanner.Text())
	}

	//Calculate col width from the operators row
	operators := pstrings[len(pstrings)-1]
	colWidths := make([]int, 0, 100)

	currWidth := 0
	col := 0
	for i := 0; i < len(operators); i++ {
		currWidth++
		if (i != 0 && operators[i] != ' ') || i == len(operators)-1 {

			if i == len(operators)-1 {
				currWidth += 2
			}

			colWidths = append(colWidths, currWidth-2)
			col++
			currWidth = 1
		}

	}

	numCols := len(colWidths)

	fmt.Println(colWidths)

	cols := make([][]string, numCols)

	// collect row values into their columns
	for _, currentString := range pstrings {
		currentCol := 0
		colWidthWithSep := 0
		for charIndex := 0; charIndex < len(currentString); charIndex += colWidthWithSep {

			colWidth := colWidths[currentCol]
			colWidthWithSep = colWidth + 1
			if currentCol == numCols-1 {
				colWidthWithSep -= 1
			}

			cols[currentCol] = append(cols[currentCol], currentString[charIndex:charIndex+colWidth])
			currentCol++
		}
	}

	finalCols := make([][]string, numCols)
	// Parse the values into their vertical number
	for c, col := range cols {
		if finalCols[c] == nil {
			finalCols[c] = make([]string, len(col))
		}

		colWidth := colWidths[c]

		// iterate over params
		for i, row := range col {

			if i != len(col)-1 {
				//Iterate over characters in param
				for j := 0; j < colWidth; j++ {
					finalCols[c][j] += string(row[j])
				}
			} else {
				finalCols[c][len(col)-1] = strings.TrimSpace(row)
			}
		}
	}

	// calculate and total the problems
	total := 0
	//loop over the problems
	for _, prob := range finalCols {
		amt := 0

		operator := prob[len(prob)-1]

		if operator == "*" {
			amt = 1
		}

		for i, param := range prob {
			if i != len(prob)-1 {
				trimmed := strings.TrimSpace(param)
				if trimmed != "" {
					parsed, err := strconv.Atoi(strings.TrimSpace(param))
					if err != nil {
						log.Fatalf("Failed to parse: %s, %s", param, err)
					}

					fmt.Printf("%d", parsed)
					if operator == "*" {
						amt *= parsed
					} else {
						amt += parsed
					}
					if i != len(prob)-1 {
						fmt.Printf(" %s ", operator)
					}
				}
			}
		}

		fmt.Printf(" = %d\n", amt)
		total += amt

	}

	fmt.Printf("Total: %d\n", total)
}

func main() {
	main_part_2_attempt_2()
}
