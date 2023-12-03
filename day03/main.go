package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var longline, newline string
var length int
var height int
var mapNumbersPos map[int]int

type coord struct {
	x int
	y int
}

var around [8]coord

func main() {
	longline = ""
	newline = ""
	length = 0
	height = 0
	around[0] = coord{1, -1}
	around[1] = coord{0, -1}
	around[2] = coord{-1, -1}
	around[3] = coord{-1, 0}
	around[4] = coord{-1, 1}
	around[5] = coord{0, 1}
	around[6] = coord{1, 1}
	around[7] = coord{1, 0}
	mapNumbersPos = make(map[int]int)

	readFile, err := os.Open("day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
		height++
	}
	length = len(fileLines[0])
	fmt.Println(length, height)
	for _, line := range fileLines {
		longline += line
		newline += line
	}
	fmt.Println(len(longline))
	{ //Part 1
		total := 0
		for y := 0; y < height; y++ {
			for x := 0; x < length; x++ {
				letra := get(longline, x, y)
				if unicode.IsDigit(rune(letra[0])) {
					current, masX, hasPart := process(longline, x, y)
					currentVal, _ := strconv.Atoi(current)
					setMap(currentVal, x, y, masX)
					if hasPart {
						total += currentVal
					}
					x += masX
				}
			}
		}
		fmt.Println("total part 1:", total, "\n\n") //514969
	}
	{ //part 2
		total := 0
		removeNonDoubleGears(newline)
		removeNonGearOperations(newline)
		removeNonGears()
		for y := 0; y < height; y++ {
			for x := 0; x < length; x++ {
				chara := get(newline, x, y)
				if chara == "*" {
					pos1 := findPos(newline, x, y, true)
					pos2 := findPos(newline, x, y, false)
					val1 := mapNumbersPos[pos1]
					val2 := mapNumbersPos[pos2]
					total += val1 * val2
					fmt.Println("gear at\t", x, y, "\t=\t", pos1, pos2, "\tvals:\t", val1, val2, "\t=>\t ", val1*val2, "\ttotal:", total)
				}
			}
		}
		//for i, chara := range newline {
		//	if i%length == 0 {
		//		fmt.Println("\t", (i/length)-1, "\t")
		//	}
		//	fmt.Print(string(chara))
		//}
		//fmt.Println("\n", mapNumbersPos)
		//43380668 wrong
		fmt.Println("total part 2:", total) //78915902 is right
	}
}

func findPos(lines string, x int, y int, clockwise bool) int {
	largo := len(around) - 1
	for i := range around {
		pos := around[i]
		if !clockwise {
			pos = around[largo-i]
		}
		xx := x + pos.x
		yy := y + pos.y
		if isDigitAt(lines, xx, yy) {
			return xx + length*yy
		}
	}
	return -1
}

func setMap(val int, x int, y int, digits int) {
	for i := 0; i < digits; i++ {
		mapNumbersPos[x+y*length+i] = val
	}
}

func removeNonDoubleGears(lines string) {
	for y := 0; y < height; y++ {
		for x := 0; x < length; x++ {
			letra := get(lines, x, y)
			if letra[0] == '*' {
				numbersAround := countNumbersAround(lines, x, y)
				fmt.Println("candidate for removal: ", x, y, "around:", numbersAround)
				if numbersAround == 2 && checksFakeDoubles(lines, x, y) {
					replaceNewLines(x, y, 1)
					fmt.Println("removing fake double gear at", x, y)
				} else if numbersAround == 3 && checksFakeTriples(lines, x, y) {
					replaceNewLines(x, y, 1)
					fmt.Println("removing fake triple gear at", x, y)
				} else if numbersAround <= 1 {
					replaceNewLines(x, y, 1)
					fmt.Println("removing single gear at", x, y)
				}
			}
		}
	}
}

func checksFakeTriples(lines string, x int, y int) bool {
	if isDigitAt(lines, x-1, y-1) && isDigitAt(lines, x, y-1) && isDigitAt(lines, x+1, y-1) {
		return true
	}
	if isDigitAt(lines, x-1, y+1) && isDigitAt(lines, x, y+1) && isDigitAt(lines, x+1, y+1) {
		return true
	}
	return false
}

func checksFakeDoubles(lines string, x int, y int) bool {
	if isDigitAt(lines, x-1, y-1) && isDigitAt(lines, x, y-1) {
		return true
	}
	if isDigitAt(lines, x, y-1) && isDigitAt(lines, x+1, y-1) {
		return true
	}
	if isDigitAt(lines, x-1, y+1) && isDigitAt(lines, x, y+1) {
		return true
	}
	if isDigitAt(lines, x, y+1) && isDigitAt(lines, x+1, y+1) {
		return true
	}
	return false
}

func countNumbersAround(lines string, x int, y int) int {
	total := 0
	for _, pos := range around {
		xx := x + pos.x
		yy := y + pos.y
		if isDigitAt(lines, xx, yy) {
			total++
		}
	}
	return total
}

func removeNonGearOperations(lines string) {
	for y := 0; y < height; y++ {
		for x := 0; x < length; x++ {
			letra := get(lines, x, y)
			if unicode.IsDigit(rune(letra[0])) {
				_, masX, hasPart := checkGears(lines, x, y)
				if !hasPart {
					replaceNewLines(x, y, masX)
				}
				x += masX
			}
		}
	}
}

func removeNonGears() {
	runeline := []rune(newline)
	for i, chara := range runeline {
		if !unicode.IsDigit(chara) && chara != '*' {
			runeline[i] = '.'
		}
	}
	newline = string(runeline)
}

func replaceNewLines(x int, y int, howMany int) {
	runeline := []rune(newline)
	newline = string(runeline[:x+y*length])
	newline += dots(howMany)
	newline += string(runeline[x+y*length+howMany:])
}

func dots(many int) string {
	val := ""
	for i := 0; i < many; i++ {
		val += "."
	}
	return val
}

func checkGears(lines string, x int, y int) (string, int, bool) {
	if !isDigitAt(lines, x, y) {
		return "", 0, false
	}
	hasPart := false
	letra := get(lines, x, y)
	for _, pos := range around {
		xx := x + pos.x
		yy := y + pos.y
		check := get(lines, xx, yy)
		if check == "*" {
			hasPart = true
		}
	}
	nextChar, distance, eitherHasPart := checkGears(lines, x+1, y)
	return letra + nextChar, distance + 1, hasPart || eitherHasPart
}

func process(lines string, x int, y int) (string, int, bool) {
	if !isDigitAt(lines, x, y) {
		return "", 0, false
	}
	hasPart := false
	letra := get(lines, x, y)
	for _, pos := range around {
		xx := x + pos.x
		yy := y + pos.y
		check := get(lines, xx, yy)
		if check != "." && !isDigitAt(lines, xx, yy) {
			hasPart = true
		}
	}
	nextChar, distance, eitherHasPart := process(lines, x+1, y)
	return letra + nextChar, distance + 1, hasPart || eitherHasPart
}

func isDigitAt(lines string, x int, y int) bool {
	return unicode.IsDigit(rune(get(lines, x, y)[0]))
}

func get(lines string, x int, y int) string {
	if x < 0 || y < 0 || x >= length || y >= height {
		return "."
	}
	pos := x + y*length
	return lines[pos : pos+1]
}
