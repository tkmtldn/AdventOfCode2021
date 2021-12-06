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

func prepareData(reader io.Reader) ([]int, int) {
	tmp := []string{}
	data := make([]int, 9)

	fileScanner := bufio.NewScanner(reader)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tmp = strings.Split(line, ",")
	}

	for _, raw := range tmp {
		v, _ := strconv.Atoi(raw)
		data[v] += 1
	}
	return data, len(tmp)
}

func lanternfish(inp []int, a int, days int) int {
	addition := 0
	for x := 0; x < days; x++ {
		new_inp := make([]int, 9)
		for k, v := range inp {
			if k == 0 {
				new_inp[8] += v
				new_inp[6] += v
				addition += v
			} else {
				new_inp[k-1] += v
			}
		}
		inp = new_inp
	}
	return addition + a
}

func main() {
	path := filepath.Join(".", "day_06", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	data, i := prepareData(fileHandle)

	fmt.Printf("First answer: %v \n", lanternfish(data, i, 80))
	fmt.Printf("Second answer: %v \n", lanternfish(data, i, 256))
}
