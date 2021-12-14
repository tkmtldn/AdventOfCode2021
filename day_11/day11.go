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

func getNeighbourhood(data [][]int, y int, x int) [][]int {
	xMax := len(data[0])
	yMax := len(data)
	adjLoc := [][]int{}

	if x-1 >= 0 {
		res := []int{y, x - 1}
		adjLoc = append(adjLoc, res)
	}
	if x+1 < xMax {
		res := []int{y, x + 1}
		adjLoc = append(adjLoc, res)
	}

	if y-1 >= 0 {
		res := []int{y - 1, x}
		adjLoc = append(adjLoc, res)
	}
	if y+1 < yMax {
		res := []int{y + 1, x}
		adjLoc = append(adjLoc, res)
	}
	if x-1 >= 0 && y-1 >= 0 {
		res := []int{y - 1, x - 1}
		adjLoc = append(adjLoc, res)
	}
	if x-1 >= 0 && y+1 < yMax {
		res := []int{y + 1, x - 1}
		adjLoc = append(adjLoc, res)
	}
	if x+1 < xMax && y-1 >= 0 {
		res := []int{y - 1, x + 1}
		adjLoc = append(adjLoc, res)
	}
	if x+1 < xMax && y+1 < yMax {
		res := []int{y + 1, x + 1}
		adjLoc = append(adjLoc, res)
	}
	return adjLoc
}

func task(data [][]int) (ans1 int, ans2 int) {
	flashed := 0
LOOP:
	for x := 1; x < 1000; x++ {
		cur_flashed := 0
		for k1, v := range data {
			for k2, _ := range v {
				data[k1][k2]++
			}
		}
		for {
			neighbours := [][]int{}
			for k1, v := range data {
				for k2, _ := range v {
					if data[k1][k2] > 9 {
						neighbours = append(neighbours, getNeighbourhood(data, k1, k2)...)
						data[k1][k2] = 0
						cur_flashed++
					}
				}
			}
			if len(neighbours) == 0 {
				break
			}
			for _, v := range neighbours {
				if data[v[0]][v[1]] == 0 {
					continue
				} else {
					data[v[0]][v[1]]++
				}
			}
		}
		flashed += cur_flashed
		if x == 100 {
			ans1 = flashed
		}
		if cur_flashed == 100 {
			ans2 = x
			break LOOP
		}
	}
	return ans1, ans2
}

func getData(fileHandle io.Reader) [][]int {
	fileScanner := bufio.NewScanner(fileHandle)
	data := [][]int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		l := strings.Split(line, "")
		newIntLine := []int{}
		for _, v := range l {
			newIntElem, _ := strconv.Atoi(v)
			newIntLine = append(newIntLine, newIntElem)
		}
		data = append(data, newIntLine)
	}
	return data
}

func main() {
	path := filepath.Join(".", "day_11", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	data := getData(fileHandle)

	ans1, ans2 := task(data)

	fmt.Printf("First answer: %v \n", ans1)
	fmt.Printf("Second answer: %v \n", ans2)
}
