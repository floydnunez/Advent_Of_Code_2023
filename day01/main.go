package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	fmt.Println(digits)

	readFile, err := os.Open("day01/input.txt")
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

	{
		total := 0
		for _, line := range fileLines {
			firstNumber := getFirstNumber(line)
			secondNumber := getLastNumber(line)
			total += (firstNumber * 10) + secondNumber
		}

		fmt.Println("part 1:", total)
	}
	{
		total := 0
		for _, line := range fileLines {
			firstNumber := getFirstNumberOrText(line, digits)
			secondNumber := getLastNumberOrText(line, digits)
			subtotal := (firstNumber * 10) + secondNumber
			fmt.Println(subtotal)
			total += subtotal
		}
		fmt.Println("part 2:", total)
	}
}

func getLastNumberOrText(line string, digits [9]string) int {
	for i := len(line) - 1; i >= 0; i-- {
		letter := line[i]
		val, done := getCharValue(letter)
		if done {
			return val
		} else {
			rest := line[i:]
			for whichDigit := 0; whichDigit < len(digits); whichDigit++ {
				if strings.HasPrefix(rest, digits[whichDigit]) {
					return whichDigit + 1
				}
			}
		}
	}
	return 0
}

func getFirstNumberOrText(line string, digits [9]string) int {
	for i := 0; i < len(line); i++ {
		letter := line[i]
		val, done := getCharValue(letter)
		if done {
			return val
		} else {
			rest := line[i:]
			for whichDigit := 0; whichDigit < len(digits); whichDigit++ {
				if strings.HasPrefix(rest, digits[whichDigit]) {
					return whichDigit + 1
				}
			}
		}
	}
	return 0
}

func getFirstNumber(line string) int {
	for i := 0; i < len(line); i++ {
		letter := line[i]
		val, done := getCharValue(letter)
		if done {
			return val
		}
	}
	return 0
}

func getCharValue(letter uint8) (int, bool) {
	if letter >= '0' && letter <= '9' {
		return int(letter - '0'), true
	}
	return 0, false
}
func getLastNumber(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		letter := line[i]
		val, done := getCharValue(letter)
		if done {
			return val
		}
	}
	return 0
}
