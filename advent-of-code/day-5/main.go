package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main_part_1() {

	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("Failed to open file with error %s", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	ranges := make([][]string, 0, 10)
	ids := make([]string, 0, 10)

	for scanner.Scan() {
		l := scanner.Text()

		t := strings.Split(l, "-")

		if len(t) == 2 {
			ranges = append(ranges, t)
		} else if len(t) == 1 {
			if len(t[0]) > 0 {
				ids = append(ids, t[0])
			}
		}
	}

	fmt.Println(ranges)
	fmt.Println(ids)

	count := 0

	for _, v := range ids {

		id, err := strconv.Atoi(v)

		if err != nil {
			log.Fatalf("Failed to parse id value: %s", v)
		}

		for _, k := range ranges {
			start, err := strconv.Atoi(k[0])
			if err != nil {
				log.Fatalf("Failed to parse range start: %s", k[0])
			}

			end, err := strconv.Atoi(k[1])
			if err != nil {
				log.Fatalf("Failed to parse range end: %s", k[1])
			}

			if id >= start && id <= end {
				count++
				break
			}

		}

	}

	fmt.Printf("Fresh ingredients: %d\n", count)

}

func main_part_2_attempt_1() {

	f, err := os.Open("sample")

	if err != nil {
		log.Fatalf("Failed to open file with error %s", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	ranges := make([][]int64, 0, 10)

	for scanner.Scan() {
		l := scanner.Text()

		t := strings.Split(l, "-")

		if len(t) == 2 {
			s, err := strconv.ParseInt(t[0], 10, 64)
			if err != nil {
				log.Fatalf("Failed to parse start %s", t[0])
			}
			e, err := strconv.ParseInt(t[1], 10, 64)
			if err != nil {
				log.Fatalf("Failed to parse end %s", t[1])
			}

			if e < s {
				log.Fatalf("Invalid range: %s", l)
			}

			ranges = append(ranges, []int64{s, e})
		}
	}

	rangeMap := make(map[string]map[string]int64)

	for i, v := range ranges {
		pk := fmt.Sprintf("%d-%d", v[0], v[1])
		rangeMap[pk] = make(map[string]int64)
		rangeMap[pk][pk] = v[1] - v[0] + 1

		for j, k := range ranges {
			if j != i {
				sk := fmt.Sprintf("%d-%d", k[0], k[1])

				_, exists := rangeMap[sk][pk]

				if !exists {
					// if any overlap
					if (v[0] <= k[1] && v[1] >= k[0]) || v[1] <= k[1] && v[1] >= k[0] {
						fmt.Printf("%d <= %d && %d >= %d || %d <= %d && %d >= %d\n", v[0], k[1], v[1], k[0], v[1], k[1], v[1], k[0])

						ss := v[0]
						se := v[1]

						if k[0] > v[0] {
							ss = k[0]
						}

						if k[1] < v[1] {

							se = k[1]
						}

						fmt.Printf("Subtracting: %d, %d = %d\n", se, se, se-ss+1)
						// count += end - start
						rangeMap[pk][sk] = se - ss + 1
					}
				}
			}
		}

	}

	count := int64(0)
	for k, v := range rangeMap {

		for sk, sv := range v {
			if sk == k {
				count += sv
			} else {
				count -= sv
			}
		}

	}

	pretty, err := json.MarshalIndent(rangeMap, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(pretty))

	fmt.Printf("Fresh ingredients: %d\n", count)

}

func main_part_2_attempt_2() {

	f, err := os.Open("input")

	if err != nil {
		log.Fatalf("Failed to open file with error %s", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	ranges := make([][]int64, 0, 10)

	for scanner.Scan() {
		l := scanner.Text()

		t := strings.Split(l, "-")

		if len(t) == 2 {
			s, err := strconv.ParseInt(t[0], 10, 64)
			if err != nil {
				log.Fatalf("Failed to parse start %s", t[0])
			}
			e, err := strconv.ParseInt(t[1], 10, 64)
			if err != nil {
				log.Fatalf("Failed to parse end %s", t[1])
			}

			if e < s {
				log.Fatalf("Invalid range: %s", l)
			}

			ranges = append(ranges, []int64{s, e})
		}
	}

	fmt.Printf("Total lines: %d\n", len(ranges))

	slices.SortFunc(ranges, func(a, b []int64) int {
		if a[0] < b[0] {
			return -1
		}

		if a[0] > b[0] {
			return 1
		}

		return 0
	})

	resets := 0
	iterations := 0
	for i := 0; i < len(ranges); i++ {

		for j := 0; j < len(ranges); j++ {
			iterations++
			if j != i {
				if ranges[i][0] <= ranges[j][1] && ranges[i][1] >= ranges[j][0] {

					if ranges[j][0] < ranges[i][0] {
						ranges[i][0] = ranges[j][0]
					}

					if ranges[j][1] > ranges[i][1] {
						ranges[i][1] = ranges[j][1]
					}

					ranges = slices.Delete(ranges, j, j+1)

					j = 0
					resets++

				}
			}
		}

	}

	count := int64(0)
	for _, v := range ranges {
		count += v[1] - v[0] + 1
	}

	//Fresh ingredients: 343143696885053, total resets: 91, iterations: 446413
	// 446413
	// 18124
	fmt.Printf("Fresh ingredients: %d, total resets: %d, iterations: %d\n", count, resets, iterations)

}

func main() {
	main_part_2_attempt_2()
}
