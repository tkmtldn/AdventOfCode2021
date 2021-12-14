package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

//I figured out part1, but stuck on part2 (too much RAM needed)
//so thanks again, Pekka (github.com/pvainio/), I made refactoring
//now it is much faster

func IsLower(s string) bool {
	return !unicode.IsUpper([]rune(s)[0])
}

func createPath(g map[string][]string, curPath []string, visited map[string]byte, maxVis byte) [][]string {
	curCave := curPath[len(curPath)-1]

	if curCave == "end" {
		return [][]string{curPath}
	}

	visited[curCave]++

	paths := [][]string{}
	for _, nextCave := range g[curCave] {
		if !alreadyMaxVisits(visited, nextCave, maxVis) {
			nextPaths := createPath(g, append(curPath, nextCave), createNew(visited), maxVis)
			paths = append(paths, nextPaths...)
		}
	}
	return paths
}

func createNew(m map[string]byte) map[string]byte {
	r := make(map[string]byte)
	for k, v := range m {
		r[k] = v
	}
	return r
}

func alreadyMaxVisits(visited map[string]byte, cave string, maxVis byte) bool {
	if !IsLower(cave) || visited[cave] == 0 {
		return false
	} else if cave == "start" || cave == "end" || maxVis < 2 {
		return true
	}

	for k, v := range visited {
		if IsLower(k) && v >= maxVis {
			return true
		}
	}

	return false
}

func addLink(links map[string][]string, a string, b string) {
	links[a] = append(links[a], b)
}

func main() {
	path := filepath.Join(".", "day_12", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	graphs := map[string][]string{}

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		link := strings.Split(line, "-")
		addLink(graphs, link[0], link[1])
		addLink(graphs, link[1], link[0])
	}

	paths1 := createPath(graphs, []string{"start"}, make(map[string]byte), 1)
	paths2 := createPath(graphs, []string{"start"}, make(map[string]byte), 2)

	fmt.Printf("First answer: %v \n", len(paths1))
	fmt.Printf("Second answer: %v \n", len(paths2))
}
