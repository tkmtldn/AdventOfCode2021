package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func addToQue(q []int, x int) int {
	q[0], q[1], q[2] = x, q[0], q[1]
	return q[0] + q[1] + q[2]
}

func main() {
	path := filepath.Join(".", "day_01", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	queue := []int{0, 0, 0}
	previous1, previous2 := 0, 0
	ans1 := 0
	ans2 := 0

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		cur, _ := strconv.Atoi(fileScanner.Text())

		if cur > previous1 {
			ans1++
		}

		curB := addToQue(queue, cur)
		if curB > previous2 {
			ans2++
		}

		previous1 = cur
		previous2 = curB
	}

	fmt.Printf("First answer: : %v \n", ans1 - 1)
	fmt.Printf("Second answer: : %v \n", ans2 - 3)
}
