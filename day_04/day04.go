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

type board struct {
	Horizontal [][]string
	Vertical   [][]string
}

func prepareData(reader io.Reader) ([]string, []board) {
	data := []string{}

	fileScanner := bufio.NewScanner(reader)
	for fileScanner.Scan() {
		data = append(data, fileScanner.Text())
	}

	bingoBalls := []string{}
	boards := []board{}

	rowsH := [][]string{}
	rowsV := [][]string{}

	for k, v := range data {
		if k == 0 {
			bingoBalls = strings.Split(v, ",")
		} else {
			if len(v) == 0 {
				continue
			} else {
				d := strings.TrimLeft(v, " ")
				d = strings.Replace(d, "  ", " ", -1)
				d_t := strings.Split(d, " ")
				rowsH = append(rowsH, d_t)
				if len(rowsH) == 5 {
					h1, h2, h3, h4, h5 := []string{}, []string{}, []string{}, []string{}, []string{}
					for y := 0; y < 5; y++ {
						h1 = append(h1, rowsH[y][0])
						h2 = append(h2, rowsH[y][1])
						h3 = append(h3, rowsH[y][2])
						h4 = append(h4, rowsH[y][3])
						h5 = append(h5, rowsH[y][4])
					}
					rowsV = append(rowsV, h1)
					rowsV = append(rowsV, h2)
					rowsV = append(rowsV, h3)
					rowsV = append(rowsV, h4)
					rowsV = append(rowsV, h5)
					newBoard := board{rowsH, rowsV}
					boards = append(boards, newBoard)
					rowsH = [][]string{}
					rowsV = [][]string{}
				}
			}
		}
	}
	return bingoBalls, boards
}

func playGame(balls []string, brds []board) ([]int, map[int][]string) {

	winners := map[int][]string{}
	order := []int{}

	for k, _ := range balls {
		ord, wBalls := playRound(balls[0:k+1], brds)

		for _, v := range ord {
			if _, ok := winners[v]; !ok {
				winners[v] = wBalls[v]
				order = append(order, v)
			}
		}
	}

	return order, winners
}

func playRound(balls []string, boards []board) ([]int, map[int][]string) {
	order := []int{}
	winBalls := map[int][]string{}

	for index, b := range boards {
		result := playRow(balls, b.Vertical, b.Horizontal)
		if result == true {
			order = append(order, index)
			winBalls[index] = balls
		}
	}
	return order, winBalls
}

func playRow(bingoBalls []string, vert [][]string, hor [][]string) bool {
	b1, b2 := false, false
	for _, v := range vert {
		res := 0
		for _, v1 := range v {
			for _, v2 := range bingoBalls {
				if v1 == v2 {
					res++
				}
			}
		}
		if res == 5 {
			b1 = true
		}
	}
	for _, v := range hor {
		res := 0
		for _, v1 := range v {
			for _, v2 := range bingoBalls {
				if v1 == v2 {
					res++
				}
			}
		}
		if res == 5 {
			b2 = true
		}
	}
	return b1 || b2
}

func findInterception(winBoard []string, playingBalls []string) (interception []string) {
	for _, v1 := range winBoard {
		for _, v2 := range playingBalls {
			if v1 == v2 {
				interception = append(interception, v2)
			}
		}
	}
	return
}

func summa(input []string) (sum int) {
	for _, v := range input {
		c, _ := strconv.Atoi(v)
		sum += c
	}
	return
}

func concat(input [][]string) (res []string) {
	for _, v := range input {
		for _, v2 := range v {
			res = append(res, v2)
		}
	}
	return
}

func main() {
	path := filepath.Join(".", "day_04", "input.txt")
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()

	bingoBalls, b := prepareData(fileHandle)

	order, winners := playGame(bingoBalls, b)

	firstWinnerBoardNum := order[0]
	lastWinnerBoardNum := order[len(order)-1]

	firstWinner := concat(b[firstWinnerBoardNum].Horizontal)
	lastWinner := concat(b[lastWinnerBoardNum].Horizontal)

	interception1 := findInterception(firstWinner, winners[firstWinnerBoardNum])
	interception2 := findInterception(lastWinner, winners[lastWinnerBoardNum])

	winningBallOne := winners[firstWinnerBoardNum][len(winners[firstWinnerBoardNum])-1]
	winningBallTwo := winners[lastWinnerBoardNum][len(winners[lastWinnerBoardNum])-1]
	winnOneInt, _ := strconv.Atoi(winningBallOne)
	winnTwoInt, _ := strconv.Atoi(winningBallTwo)

	res1 := winnOneInt * (summa(firstWinner) - summa(interception1))
	res2 := winnTwoInt * (summa(lastWinner) - summa(interception2))

	fmt.Printf("First answer: %v\n", res1)
	fmt.Printf("Second answer: %v\n", res2)
}
