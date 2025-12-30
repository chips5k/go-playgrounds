package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main_part_1() {

	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("Faile to open file with error '%s'", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make([][]string, 0, 10)

	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, strings.Split(l, ""))
	}

	count := 0
	for i := range grid {

		for j := range grid[i] {
			proximityCount := 0

			if grid[i][j] != "." {
				// If prev col is not blank
				if j-1 >= 0 {
					if grid[i][j-1] != "." {
						proximityCount++
					}
				}

				// if next col is not blank
				if j+1 < len(grid[i]) {
					if grid[i][j+1] != "." {
						proximityCount++
					}
				}

				// if prev row exists
				if i-1 >= 0 {
					// If prev col is not blank
					if j-1 >= 0 {
						if grid[i-1][j-1] != "." {
							proximityCount++
						}
					}

					// if mid col is not blank
					if grid[i-1][j] != "." {
						proximityCount++
					}

					// if next col is not blank
					if j+1 < len(grid[i-1]) {
						if grid[i-1][j+1] != "." {
							proximityCount++
						}
					}
				}

				// if next row exists
				if i+1 < len(grid) {
					// If prev col is not blank
					if j-1 >= 0 {
						if grid[i+1][j-1] != "." {
							proximityCount++
						}
					}

					// if mid col is not blank
					if grid[i+1][j] != "." {
						proximityCount++
					}

					// if next col is not blank
					if j+1 < len(grid[i+1]) {
						if grid[i+1][j+1] != "." {
							proximityCount++
						}
					}
				}

				if proximityCount < 4 {
					count++
				}
			}

		}

	}

	fmt.Println(count)

}

func main_part)2() {

	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("Faile to open file with error '%s'", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make([][]string, 0, 10)

	for scanner.Scan() {
		l := scanner.Text()
		grid = append(grid, strings.Split(l, ""))
	}

	count := 0
	fmt.Println(grid)
	fmt.Println()
	removed := true
	for removed != false {
		totalRemoved := 0
		for i := range grid {

			for j := range grid[i] {
				proximityCount := 0

				if grid[i][j] != "." {
					// If prev col is not blank
					if j-1 >= 0 {
						if grid[i][j-1] != "." {
							proximityCount++
						}
					}

					// if next col is not blank
					if j+1 < len(grid[i]) {
						if grid[i][j+1] != "." {
							proximityCount++
						}
					}

					// if prev row exists
					if i-1 >= 0 {
						// If prev col is not blank
						if j-1 >= 0 {
							if grid[i-1][j-1] != "." {
								proximityCount++
							}
						}

						// if mid col is not blank
						if grid[i-1][j] != "." {
							proximityCount++
						}

						// if next col is not blank
						if j+1 < len(grid[i-1]) {
							if grid[i-1][j+1] != "." {
								proximityCount++
							}
						}
					}

					// if next row exists
					if i+1 < len(grid) {
						// If prev col is not blank
						if j-1 >= 0 {
							if grid[i+1][j-1] != "." {
								proximityCount++
							}
						}

						// if mid col is not blank
						if grid[i+1][j] != "." {
							proximityCount++
						}

						// if next col is not blank
						if j+1 < len(grid[i+1]) {
							if grid[i+1][j+1] != "." {
								proximityCount++
							}
						}
					}

					if proximityCount < 4 {
						totalRemoved++
						grid[i][j] = "."
					}
				}

			}

		}

		if totalRemoved > 0 {
			removed = true
		} else {
			removed = false
		}
		count += totalRemoved
	}

	fmt.Println(grid)
	fmt.Println(count)

}

func main() {
	main_part_2()
}
