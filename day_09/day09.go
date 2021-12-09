package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func getNeighbourhood(data [][]string, y int, x int) [][]int {
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
	return adjLoc
}

func checkLowest(data [][]string, y int, x int) bool {
	adjLoc := getNeighbourhood(data, y, x)

	s := 4 - len(adjLoc)
	for _, v := range adjLoc {
		if data[v[0]][v[1]] > data[y][x] {
			s++
		}
	}
	return s == 4
}

func sumLowPoints(data [][]string) (int, [][]int) {
	sum := 0
	lowestPoints := [][]int{}
	for k1, v1 := range data {
		for k2, _ := range (v1) {
			ans := checkLowest(data, k1, k2)
			if ans == true {
				i, _ := strconv.Atoi(data[k1][k2])
				sum += i + 1
				res := []int{k1, k2}
				lowestPoints = append(lowestPoints, res)
			}
		}
	}
	return sum, lowestPoints
}

func getThreeLargestBasins(data [][]string, lows [][]int) int {
	basins := []int{}
	for _, v := range lows {
		basinSize := countBasinSize(data, v)
		basins = append(basins, basinSize)
	}
	sort.Ints(basins)
	return basins[len(basins)-1]*basins[len(basins)-2]*basins[len(basins)-3]
}

func countBasinSize(data [][]string, work []int) int {
	n := [][]int{}
	workingSet := map[string]int{}

	n = append(n, work)
	name := string(work[0]) + "-" + string(work[1])
	workingSet[name] = 0

	for {
		startLen := len(workingSet)
		new:=[][]int{}
		for _, v := range n {
			neigb := getNeighbourhood(data, v[0], v[1])
			for _, value := range neigb {
				if data[value[0]][value[1]] != "9" {
					new = append(new, value)
				}
			}
		}
		for _, v := range new{
			name := string(v[0]) + "-" + string(v[1])
			workingSet[name] = 0
			n = append(n, v)
		}

		finishLen := len(workingSet)
		if startLen == finishLen {
			break
		}
	}

	return len(workingSet)
}

func main() {
	path := filepath.Join(".", "day_09", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	data := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		l := strings.Split(line, "")
		data = append(data, l)
	}

	sum1, lowestPoints := sumLowPoints(data)
	sum2 := getThreeLargestBasins(data, lowestPoints)

	fmt.Printf("First answer: %v \n", sum1)
	fmt.Printf("Second answer: %v \n", sum2)
}
