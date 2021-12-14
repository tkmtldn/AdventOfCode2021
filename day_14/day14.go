package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var existingCounts = make(map[string]map[byte]int64)
var pairs = make(map[string]byte)

// Thank you for your help, Pekka
// https://github.com/pvainio/adventofcode/blob/main/2021/go/d14/main.go

func prepareData(fileHandle io.Reader) string {
	fileScanner := bufio.NewScanner(fileHandle)
	template := ""

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, " -> "){
			v := strings.Split(fileScanner.Text(), " -> ")
			pairs[v[0]] = v[1][0]
		} else if line != "" {
			template = line
		}
	}
	return template
}

func coutnPolymer(t string, n int) int64 {

	result := make(map[byte]int64)
	pairs := stringAsCharPairs(t) //Хранит комбинации 01,12,23 итд

	for pos, pair := range pairs {
		lastPair := pos == len(pairs)-1
		charCountsForPair := countCharsForPair(pair, lastPair, n)
		result = mergeCountsByChar(result, charCountsForPair)
	}

	var min int64 = math.MaxInt64
	var max int64 = 0
	for _, v := range result {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func countCharsForPair(pair []byte, lastPair bool, depth int) map[byte]int64 {
	if depth == 0 {
		if lastPair {
			return map[byte]int64{pair[0]: 1, pair[1]: 1}
		} else {
			return map[byte]int64{pair[0]: 1}
		}
	}
	key := fmt.Sprintf("%v%v%v", string(pair), lastPair, depth)
	if res, ok := existingCounts[key]; ok {
		return res
	} else {
		pairToChar := pairs[string(pair)]
		pair1 := []byte{pair[0], pairToChar}
		pair2 := []byte{pairToChar, pair[1]}
		p1Count := countCharsForPair(pair1, false, depth-1)
		p2Count := countCharsForPair(pair2, lastPair, depth-1)
		res := mergeCountsByChar(p1Count, p2Count)
		existingCounts[key] = res
		return res
	}
}

func mergeCountsByChar(a map[byte]int64, b map[byte]int64) map[byte]int64 {
	res := make(map[byte]int64)
	addCounts(res, a)
	addCounts(res, b)
	return res
}

func addCounts(res map[byte]int64, source map[byte]int64) {
	for k, v := range source {
		if existing, ok := res[k]; ok {
			res[k] = existing + v
		} else {
			res[k] = v
		}
	}
}

func stringAsCharPairs(str string) [][]byte {
	var prev byte
	var result [][]byte
	for _, c := range str {
		if prev != 0 {
			result = append(result, []byte{prev, byte(c)})
		}
		prev = byte(c)
	}
	return result
}

func main() {
	path := filepath.Join(".", "day_14", "testinput.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	template := prepareData(fileHandle)

	ans1 := coutnPolymer(template, 10)
	fmt.Printf("First answer: %v \n", ans1)

	ans2 := coutnPolymer(template, 40)
	fmt.Printf("Second answer: %v \n", ans2)
}