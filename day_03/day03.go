package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func mostAndLeastPopular(data []string) (string, string) {

	fileLength := len(data)
	lineLength := len(data[0])
	zeroesCount := make([]int, lineLength)

	for _, line := range data {
		splitted := strings.Split(line, "")
		for i, value := range splitted {
			if value == "0" {
				zeroesCount[i] += 1
			}
		}
	}

	most := ""
	least := ""
	for _, v := range zeroesCount {
		if v > fileLength/2 {
			most += "0"
			least += "1"
		} else {
			most += "1"
			least += "0"
		}
	}
	return most, least
}

func countPowerConsumption(most, least string) int64 {
	a, _ := strconv.ParseInt(most, 2, 64)
	b, _ := strconv.ParseInt(least, 2, 64)
	return a * b
}

func exitFromRating(data []string, prefix string) string {
	for _, line := range data {
		if strings.HasPrefix(line, prefix) {
			return line
		}
	}
	return ""
}

func rating(data []string, commonInput string, variant bool) string {
	num0, num1 := 0, 0
	lineLength := len(data[0])
	prefix := string(commonInput[0])

	for t := 1; t < lineLength; t++ {
		for _, line := range data {
			if strings.HasPrefix(line, prefix) {
				linear := strings.Split(line, "")
				if linear[t] == "0" {
					num0 ++
				} else {
					num1 ++
				}
			}
		}
		if variant {
			if num1 >= num0 {
				prefix += "1"
			} else {
				prefix += "0"
			}
		} else {
			if num1+num0 <= 1 {
				return exitFromRating(data, prefix)
			} else if num1 < num0 {
				prefix += "1"
			} else {
				prefix += "0"
			}
		}
		num1, num0 = 0, 0
	}
	return prefix
}

func main() {
	path := filepath.Join(".", "day_03", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	data := make([]string, 0)
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		data = append(data, line)
	}

	mostCommon, leastCommon := mostAndLeastPopular(data)
	fmt.Printf("First answer: %v \n", countPowerConsumption(mostCommon, leastCommon))

	oxygenGen := rating(data, mostCommon, true)
	CO2scrubber := rating(data, leastCommon, false)
	fmt.Printf("Second answer: %v \n", countPowerConsumption(oxygenGen, CO2scrubber))
}