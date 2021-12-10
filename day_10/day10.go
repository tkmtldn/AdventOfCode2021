package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func sumLine(s []int32) (sum int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	for _, v := range s {
		sum *= 5
		if v == 40 {
			sum += 1
		}
		if v == 60 {
			sum += 4
		}
		if v == 91 {
			sum += 2
		}
		if v == 123 {
			sum += 3
		}
	}
	return sum
}

func main() {
	path := filepath.Join(".", "day_10", "testinput.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	result1 := 0
	result2 := []int{}

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		last := 0
		cal := []int32{}
		failed := false

	LOOP:
		for _, x := range line {
			if x == 40 || x == 60 || x == 91 || x == 123 {
				last = int(x)
				cal = append(cal, x)
			}
			if x == 41 {
				//()
				if last != 40 {
					result1 += 3
					failed = true
					break LOOP
				} else {
					last = int(cal[len(cal)-2])
					cal = cal[:len(cal)-1]
				}
			}
			if x == 62 {
				//<>
				if last != 60 {
					result1 += 25137
					failed = true
					break LOOP
				} else {
					last = int(cal[len(cal)-2])
					cal = cal[:len(cal)-1]
				}
			}
			if x == 93 {
				//[]
				if last != 91 {
					result1 += 57
					failed = true
					break LOOP
				} else {
					if len(cal) < 2 {
						last = 0
						cal = []int32{}
					} else {
						last = int(cal[len(cal)-2])
						cal = cal[:len(cal)-1]
					}
				}
			}
			if x == 125 {
				//{}
				if last != 123 {
					result1 += 1197
					failed = true
					break LOOP
				} else {
					if len(cal) < 2 {
						last = 0
						cal = []int32{}
					} else {
						last = int(cal[len(cal)-2])
						cal = cal[:len(cal)-1]
					}
				}
			}
		}

		if !failed {
			result2 = append(result2, sumLine(cal))
		}
	}

	sort.Ints(result2)
	fmt.Printf("First answer: %v \n", result1)
	fmt.Printf("Second answer: %v \n", result2[(len(result2)-1)/2])
}
