package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type coordinates struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func prepareData(reader io.Reader) (output []coordinates) {
	fileScanner := bufio.NewScanner(reader)
	for fileScanner.Scan() {
		c := coordinates{}
		line := fileScanner.Text()
		line = strings.Replace(line, " -> ", ",", -1)
		coord := strings.Split(line, ",")
		c.x1, _ = strconv.Atoi(coord[0])
		c.y1, _ = strconv.Atoi(coord[1])
		c.x2, _ = strconv.Atoi(coord[2])
		c.y2, _ = strconv.Atoi(coord[3])
		output = append(output, c)
	}
	return
}

func out1(coo []coordinates) (result int) {
	problem := linearCount(coo)

	for _, v := range problem {
		if v > 0 {
			result ++
		}
	}
	return result

}

func checkInMap(problem map[string]int, x int, y int) {
	name := strconv.Itoa(x) + "-" + strconv.Itoa(y)
	if _, ok := problem[name]; !ok {
		problem[name] = 0
	} else {
		problem[name] += 1
	}
}

func linearCount(coo []coordinates) map[string]int{
	problem := map[string]int{}

	for _, c := range coo {
		if c.y1 == c.y2 {
			if c.x1 < c.x2 {
				for x := c.x1; x <= c.x2; x++ {
					checkInMap(problem, x, c.y1)
				}
			} else {
				for x := c.x1; x >= c.x2; x-- {
					checkInMap(problem, x, c.y1)
				}
			}
		}
		if c.x1 == c.x2 {
			if c.y1 < c.y2 {
				for y := c.y1; y <= c.y2; y++ {
					checkInMap(problem, c.x1, y)
				}
			} else {
				for y := c.y1; y >= c.y2; y-- {
					checkInMap(problem, c.x1, y)
				}
			}
		}
	}
	return problem
}

func out2(coo []coordinates) (result int) {
	problem := linearCount(coo)

	//diagonalCount
	for _, c := range coo {
		if math.Abs(float64(c.x1)-float64(c.x2)) == math.Abs(float64(c.y1)-float64(c.y2)) {

			if c.x1 > c.x2 && c.y1 > c.y2 {
				rang := math.Abs(float64(c.y1) - float64(c.y2))
				for i := 0; i <= int(rang); i++ {
					checkInMap(problem, c.x1-i, c.y1-i)
				}
			} else if c.x1 > c.x2 && c.y1 < c.y2 {
				rang := math.Abs(float64(c.y1) - float64(c.y2))
				for i := 0; i <= int(rang); i++ {
					checkInMap(problem, c.x1-i, c.y1+i)
				}
			} else if c.x1 < c.x2 && c.y1 > c.y2 {
				rang := math.Abs(float64(c.y1) - float64(c.y2))
				for i := 0; i <= int(rang); i++ {
					checkInMap(problem, c.x1+i, c.y1-i)
				}
			} else if c.x1 < c.x2 && c.y1 < c.y2 {
				rang := math.Abs(float64(c.y1) - float64(c.y2))
				for i := 0; i <= int(rang); i++ {
					checkInMap(problem, c.x1+i, c.y1+i)
				}
			}
		}
	}

	for _, v := range problem {
		if v > 0 {
			result ++
		}
	}
	return result
}

func main() {
	path := filepath.Join(".", "day_05", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	data := prepareData(fileHandle)

	fmt.Printf("First answer: %v \n", out1(data))
	fmt.Printf("Second answer: %v \n", out2(data))
}
