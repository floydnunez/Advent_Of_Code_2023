package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day05/input.txt")
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
	seed2soil := make(map[int]int)
	soil2fert := make(map[int]int)
	fert2water := make(map[int]int)
	water2light := make(map[int]int)
	light2temp := make(map[int]int)
	temp2humid := make(map[int]int)
	humid2loca := make(map[int]int)

	seeds := parseEach(fileLines[0], true)
	fmt.Println("seeds:", seeds)
	var currmap map[int]int
	for i := 2; i < len(fileLines); i++ {
		line := fileLines[i]
		if line == "" {
			continue
		} else if line == "seed-to-soil map:" {
			currmap = seed2soil
			fmt.Println(line, currmap)
		} else if line == "soil-to-fertilizer map:" {
			currmap = soil2fert
			fmt.Println(line, currmap)
		} else if line == "fertilizer-to-water map:" {
			currmap = fert2water
			fmt.Println(line, currmap)
		} else if line == "water-to-light map:" {
			currmap = water2light
			fmt.Println(line, currmap)
		} else if line == "light-to-temperature map:" {
			currmap = light2temp
			fmt.Println(line, currmap)
		} else if line == "temperature-to-humidity map:" {
			currmap = temp2humid
			fmt.Println(line, currmap)
		} else if line == "humidity-to-location map:" {
			currmap = humid2loca
			fmt.Println(line, currmap)
		} else {
			numbers := parseEach(line, false)
			fmt.Println("line:", line, numbers, currmap)
			for kk := 0; kk < numbers[2]; kk++ {
				currmap[numbers[1]+kk] = numbers[0] + kk
			}
		}
	}
	fmt.Println(seeds)
	fmt.Println(seed2soil)
	fmt.Println(soil2fert)
	fmt.Println(fert2water)
	fmt.Println(water2light)
	fmt.Println(light2temp)
	fmt.Println(temp2humid)
	fmt.Println(humid2loca)

	var result []int
	for _, seed := range seeds {
		soil := transform(seed2soil, seed)
		fert := transform(soil2fert, soil)
		water := transform(fert2water, fert)
		light := transform(water2light, water)
		temp := transform(light2temp, light)
		humid := transform(temp2humid, temp)
		loca := transform(humid2loca, humid)
		result = append(result, loca)
	}
	fmt.Println("result:", result)
	fmt.Println("Part 1:", slices.Min(result))
}

func transform(curmap map[int]int, key int) int {
	val, ok := curmap[key]
	if !ok {
		return key
	} else {
		return val
	}

}

func parseEach(line string, skipFirst bool) []int {
	var result []int
	parts := strings.Split(line, " ")
	i := 0
	if skipFirst {
		i = 1
	}
	for ; i < len(parts); i++ {
		part := parts[i]
		parsed, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			fmt.Println("parts:", parts, part)
			println("Error!", err)
			os.Exit(3)
		}
		result = append(result, parsed)
	}
	return result
}
