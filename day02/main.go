package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cubes struct {
	game  string
	red   int
	green int
	blue  int
}

func main() {
	readFile, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	errclose := readFile.Close()
	if errclose != nil {
		return
	}

	max := cubes{"all", 12, 13, 14}

	total1 := 0
	total2 := 0

	for _, line := range fileLines {
		game := strings.Split(line, ":")
		gameNumber := strings.Split(game[0], " ")
		fmt.Println("game number:", gameNumber[1])
		sets := trimEach(strings.Split(game[1], ";"))
		curr := cubes{gameNumber[1], 0, 0, 0}

		for _, set := range sets {
			colors := trimEach(strings.Split(set, ","))
			for _, color := range colors {
				pair := strings.Split(color, " ")
				val, _ := strconv.Atoi(pair[0])
				if strings.HasPrefix(pair[1], "r") && val > curr.red {
					curr.red = val
				}
				if strings.HasPrefix(pair[1], "g") && val > curr.green {
					curr.green = val
				}
				if strings.HasPrefix(pair[1], "b") && val > curr.blue {
					curr.blue = val
				}
			}
		}
		if curr.red <= max.red && curr.green <= max.green && curr.blue <= max.blue {
			one, _ := strconv.Atoi(curr.game)
			total1 += one
			fmt.Println("possible:", curr)
		}
		total2 += curr.red * curr.green * curr.blue
	}
	fmt.Println("part 1 total: ", total1)
	fmt.Println("part 2 total: ", total2)
}

func trimEach(sets []string) []string {
	for i := range sets {
		sets[i] = strings.TrimSpace(sets[i])
	}
	return sets
}
