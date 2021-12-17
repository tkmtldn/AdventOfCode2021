package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type area struct {
	x1, x2, y1, y2 int
}

func readFile() []int {
	path := filepath.Join(".", "day_17", "input.txt")

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	text = strings.Replace(text, ", y=", "..",1)
	text = strings.Replace(text, "target area: x=", "", 1)
	a := strings.Split(text, "..")
	ans :=[]int{}
	for _, v:= range a{
		i, _ := strconv.Atoi(v)
		ans = append(ans, i)
	}
	return ans
}

func shoot(xv int, yv int, target area) (bool, int) {
	var x, y, maxHeight int
	for {
		x, y = x+xv, y+yv
		xv, yv = max(0, xv-1), yv-1
		maxHeight = max(maxHeight, y)

		if x >= target.x1 && x <= target.x2 && y <= target.y1 && y >= target.y2 {
			return true, maxHeight
		} else if y < target.y2 || x > target.x2 {
			return false, 0
		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var target area
	data := readFile()
	target.x1 = data[0]
	target.x2 = data[1]
	target.y2 = data[2]
	target.y1 = data[3]

	maxHeight := 0
	distinctVelocityValue := 0

	for xv := 0; xv <= target.x2; xv++ {
		for yv := target.y2; yv < 1000; yv++ {
			goal, height := shoot(xv, yv, target)
			if goal {
				distinctVelocityValue++
				maxHeight = max(maxHeight, height)
			}
		}
	}

	fmt.Printf("First answer: %v \n", maxHeight)
	fmt.Printf("Second answer: %v \n", distinctVelocityValue)
}