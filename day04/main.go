package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	number  int
	winners []string
	numbers []string
	points  int
	matches int
}

func main() {
	readFile, err := os.Open("day04/input.txt")
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

	total := 0
	var initialCards []card
	for _, line := range fileLines {
		lineVal := 0
		matches := 0
		card_numbers := trimEach(strings.Split(line, ":"))
		card_number := strings.FieldsFunc(card_numbers[0], isSpace)
		number, _ := strconv.Atoi(card_number[1])
		winners_numbers := trimEach(strings.Split(card_numbers[1], "|"))
		list_winners := strings.FieldsFunc(winners_numbers[0], isSpace)
		list_numbers := strings.FieldsFunc(winners_numbers[1], isSpace)
		for _, winner := range list_winners {
			if slices.Contains(list_numbers, winner) {
				matches++
				if lineVal == 0 {
					lineVal = 1
				} else {
					lineVal = 2 * lineVal
				}
			}
		}
		oneCard := card{number: number, winners: list_winners, numbers: list_numbers,
			points: lineVal, matches: matches}
		initialCards = append(initialCards, oneCard)
		total += lineVal
	}
	fmt.Println(initialCards)
	fmt.Println("total part 1:", total) //21919, 50811 was wrong

	var listCards []card
	listCards = append(listCards, initialCards...)
	for ii := 0; ii < len(listCards); ii++ {
		current := listCards[ii]
		for kk := 1; kk <= current.matches; kk++ {
			newCard := initialCards[(current.number-1)+kk]
			listCards = append(listCards, newCard)
		}
	}
	fmt.Println("total part 2:", len(listCards)) //9881048
}

func trimEach(sets []string) []string {
	for i := range sets {
		sets[i] = strings.TrimSpace(sets[i])
	}
	return sets
}

func isSpace(c rune) bool {
	return c == ' '
}
