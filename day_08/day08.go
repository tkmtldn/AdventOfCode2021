package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func collectData(r io.Reader) ([][]string, [][]string) {
	fileScanner := bufio.NewScanner(r)
	ans0 := [][]string{}
	ans1 := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitted := strings.Split(line, " | ")
		d0 := strings.Split(splitted[0], " ")
		d1 := strings.Split(splitted[1], " ")
		ans0 = append(ans0, d0)
		ans1 = append(ans1, d1)
	}
	return ans0, ans1
}

func part1(ans [][]string) (sum int) {
	for _, v1 := range ans {
		for _, v := range v1 {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				sum++
			}
		}
	}
	return sum
}

func part2(ans0 [][]string, ans1 [][]string) (sum int) {
	for x:=0; x<len(ans0); x++{
		stringAns := []int{}
		alphabet := mapping(ans0[x])
		for _,v:= range ans1[x]{
			for key, value := range alphabet{
				if (SortStringByCharacter(v))==value{
					stringAns = append(stringAns, key)
				}
			}
		}
		sum += stringAns[0]*1000
		sum += stringAns[1]*100
		sum += stringAns[2]*10
		sum += stringAns[3]*1
	}
	return sum
}

func mapping(inp []string) (map[int]string) {
	twoThreeFive := []string{}
	zeroSixNine := []string{}
	ans := map[int]string{}
	first1, first2 := "", ""
	for _, v := range inp {
		if len(v) == 2 {
			ans[1] = v
			first1 = string(v[0])
			first2 = string(v[1])
		} else if len(v) == 3 {
			ans[7] = v
		} else if len(v) == 4 {
			ans[4] = v
		} else if len(v) == 7 {
			ans[8] = v
		} else if len(v) == 5 {
			twoThreeFive = append(twoThreeFive, v)
		} else {
			zeroSixNine = append(zeroSixNine, v)
		}
	}
	hive := []string{}
	for _, v := range twoThreeFive {
		if strings.Contains(v, first1) && strings.Contains(v, first2) {
			ans[3] = v
		} else {
			hive = append(hive, v)
		}
	}
	for _, v := range zeroSixNine {
		if !(strings.Contains(v, first1) && strings.Contains(v, first2)) {
			ans[6] = v
		} else {
			hive = append(hive, v)
		}
	}

	hivemap := map[string]int{}
	for _, v1 := range hive {
		for _, v2 := range v1 {
			if _, ok := hivemap[string(v2)]; ok {
				hivemap[string(v2)]++
			} else {
				hivemap[string(v2)] = 1
			}
		}
	}
	selectPoint := ""
	for k, v := range hivemap {
		if v == 2 {
			selectPoint = k
		}
	}

	for _, v := range hive {
		if len(v) == 5 && strings.Contains(v, selectPoint) {
			ans[2] = v
		} else if len(v) == 5 && !strings.Contains(v, selectPoint) {
			ans[5] = v
		} else if len(v) == 6 && strings.Contains(v, selectPoint) {
			ans[0] = v
		} else if len(v) == 6 && !strings.Contains(v, selectPoint) {
			ans[9] = v
		}
	}
	for key, value := range ans{
		ans[key] = SortStringByCharacter(value)
	}
	return ans
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	path := filepath.Join(".", "day_08", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	data0, data1 := collectData(fileHandle)

	fmt.Printf("First answer: %v \n", part1(data1))
	fmt.Printf("Second answer: %v \n", part2(data0, data1))
}
