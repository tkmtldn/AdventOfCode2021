package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getInput(fileHandle io.Reader) ([][]int, [][]string) {
	dots := [][]int{}
	folds := [][]string{}
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !strings.HasPrefix(line, "fold along ") && len(line) != 0 {
			lineS := strings.Split(line, ",")
			x, _ := strconv.Atoi(lineS[0])
			y, _ := strconv.Atoi(lineS[1])
			lineS2 := []int{x, y}
			dots = append(dots, lineS2)
		} else if len(line) != 0 {
			line = strings.Replace(line, "fold along ", "", -1)
			lineS := strings.Split(line, "=")
			folds = append(folds, lineS)
		}
	}
	return dots, folds
}

func dotsCount(dots [][]int, folds [][]string, second bool) int {
	dotsLen := []int{}

	max_x, max_y := 0, 0

	for _, val := range folds {
		num, _ := strconv.Atoi(val[1])
		if val[0] == "y" {
			max_y = num
			for _, v := range dots {
				if v[1] >= num {
					v[1] = num*2 - v[1]
				}
			}
		}
		if val[0] == "x" {
			for _, v := range dots {
				max_x = num
				if v[0] >= num {
					v[0] = num*2 - v[0]
				}
			}

		}
		dotsMap := map[string]int{}
		for _, v := range dots {
			name := string(v[0]) + "" + string(v[1])
			dotsMap[name] = 0
		}
		dotsLen = append(dotsLen, len(dotsMap))
	}
	dotsMap := map[string]int{}
	for _, v := range dots {
		name := string(v[0]) + "" + string(v[1])
		dotsMap[name] = 0
	}

	if !second {
		return dotsLen[0]
	} else {
		for i := 0; i < max_y; i++ {
			for j := 0; j < max_x; j++ {
				name := string(j) + "" + string(i)
				if _, ok := dotsMap[name]; ok {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
	return 0
}

func main() {
	path := filepath.Join(".", "day_13", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	dots, folds := getInput(fileHandle)

	ans1 := dotsCount(dots, folds, false)
	ans2 := dotsCount(dots, folds, true)

	fmt.Printf("First answer: %v \n", ans1)
	fmt.Printf("Second answer: %v \n", ans2)
}
