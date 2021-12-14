package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(".", "day_15", "testinput.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	fmt.Printf("First answer: %v \n", )
	fmt.Printf("Second answer: %v \n", )
}