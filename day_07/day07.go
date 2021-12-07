package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	path := filepath.Join(".", "day_07", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	data := []int{}
	for fileScanner.Scan() {
		d := fileScanner.Text()
		da := strings.Split(d, ",")
		for _, v := range da {
			dat, _ := strconv.Atoi(v)
			data = append(data, dat)
		}
	}

	_,max := MinMax(data)
	res1 := []int{}
	res2 := []int{}
	for x:=1; x<=max; x++{
		sum := 0
		sum2 := 0
		for _, v:= range data {
			rangeHere := int(math.Abs(float64(v-x)))
			sum += rangeHere
			for y:=1; y<=rangeHere; y++ {
				sum2+=y
			}
		}
		res1 = append(res1, sum)
		res2 = append(res2, sum2)
	}

	min1, _ :=MinMax(res1)
	min2, _ :=MinMax(res2)

	fmt.Printf("First answer: %v \n", min1)
	fmt.Printf("Second answer: %v \n", min2)
}
