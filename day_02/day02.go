package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path := filepath.Join(".", "day_02", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	horizontal := 0
	depth1 := 0
	depth2 := 0

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		action := strings.Split(str, " ")
		move, _ := strconv.Atoi(action[1])
		switch action[0] {
		case "forward":
			horizontal += move
			depth2 += move * depth1
		case "down":
			depth1 += move
		case "up":
			depth1 -= move
		}
	}
	fmt.Printf("First answer: %v \n", horizontal * depth1)
	fmt.Printf("Second answer: %v \n", horizontal * depth2)
}
