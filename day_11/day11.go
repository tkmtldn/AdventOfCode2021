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

func first(data [][]int) {
	for x := 0; x < 3; x++ {
		fmt.Println(data)
		// First, the energy level of each octopus increases by 1.
		for k1, v := range data {
			for k2, _ := range v {
				data[k1][k2]++
			}
		}
		// Then, any octopus with an energy level greater than 9 flashes.
		// This increases the energy level of all adjacent octopuses by 1,
		// including octopuses that are diagonally adjacent. If this causes
		// an octopus to have an energy level greater than 9, it also flashes.
		// This process continues as long as new octopuses keep having their
		// energy level increased beyond 9. (An octopus can only flash at most once per step.)
		flashedSum := 0
		for k1, v := range data {
			for k2, _ := range v {
				if data[k1][k2] > 9 {
					flashedSum++
				}
			}
		}
		// Finally, any octopus that flashed during this step has its energy
		// level set to 0, as it used all of its energy to flash.
	}
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
	path := filepath.Join(".", "day_11", "testinput.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	data := getData(fileHandle)

	first(data)

	fmt.Printf("First answer: %v \n", )
	fmt.Printf("Second answer: %v \n", )
}
