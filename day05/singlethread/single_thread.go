package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type mapping struct {
	dest   int
	source int
	length int
}
type pairs struct {
	ini int
	len int
}

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
	var seed2soil []mapping
	var soil2fert []mapping
	var fert2water []mapping
	var water2light []mapping
	var light2temp []mapping
	var temp2humid []mapping
	var humid2loca []mapping

	seeds := parseEach(fileLines[0], true)
	fmt.Println("seeds:", seeds)
	line := ""
	for i := 2; i < len(fileLines); i++ {
		currline := fileLines[i]
		if currline == "" {
			continue
		} else if strings.Contains(currline, "-to-") {
			line = currline
			continue
		}
		if line == "seed-to-soil map:" {
			numbers := parseEach(currline, false)
			seed2soil = append(seed2soil, mapping{dest: numbers[0], source: numbers[1], length: numbers[2]})
		} else if line == "soil-to-fertilizer map:" {
			numbers := parseEach(currline, false)
			soil2fert = append(soil2fert, mapping{dest: numbers[0], source: numbers[1], length: numbers[2]})
		} else if line == "fertilizer-to-water map:" {
			numbers := parseEach(currline, false)
			fert2water = append(fert2water, mapping{dest: numbers[0], source: numbers[1], length: numbers[2]})
		} else if line == "water-to-light map:" {
			numbers := parseEach(currline, false)
			water2light = append(water2light, mapping{dest: numbers[0], source: numbers[1], length: numbers[2]})
		} else if line == "light-to-temperature map:" {
			numbers := parseEach(currline, false)
			light2temp = append(light2temp, mapping{dest: numbers[0], source: numbers[1], length: numbers[2]})
		} else if line == "temperature-to-humidity map:" {
			numbers := parseEach(currline, false)
			temp2humid = append(temp2humid, mapping{dest: numbers[0], source: numbers[1], length: numbers[2]})
		} else if line == "humidity-to-location map:" {
			numbers := parseEach(currline, false)
			humid2loca = append(humid2loca, mapping{dest: numbers[0], source: numbers[1], length: numbers[2]})
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
		location := transformFull(seed, seed2soil, soil2fert, fert2water, water2light, light2temp, temp2humid, humid2loca)
		result = append(result, location)
	}
	fmt.Println("result:", result)
	fmt.Println("Part 1:", slices.Min(result)) //218513636

	//Part 2
	var seedPairs []pairs
	for ii := 0; ii < len(seeds); ii += 2 {
		sp := pairs{ini: seeds[ii], len: seeds[ii+1]}
		seedPairs = append(seedPairs, sp)
	}
	fmt.Println("seed pairs:", seedPairs)
	var locaPart2 []int
	for _, pair := range seedPairs {
		fmt.Println("processing: ", pair)
		for dynaseed := pair.ini; dynaseed < pair.ini+pair.len; dynaseed++ {
			location := transformFull(dynaseed, seed2soil, soil2fert, fert2water, water2light, light2temp, temp2humid, humid2loca)
			locaPart2 = append(locaPart2, location)
		}
	}
	fmt.Println("locapart 2 len:", len(locaPart2)) //2221837783
	fmt.Println("Part 2:", slices.Min(locaPart2))  //81956384
}

func transformFull(seed int, seed2soil []mapping, soil2fert []mapping, fert2water []mapping,
	water2light []mapping, light2temp []mapping, temp2humid []mapping, humid2loca []mapping) int {
	soil := transform(seed2soil, seed)
	fert := transform(soil2fert, soil)
	water := transform(fert2water, fert)
	light := transform(water2light, water)
	temp := transform(light2temp, light)
	humid := transform(temp2humid, temp)
	loca := transform(humid2loca, humid)
	return loca
}

func transform(currmap []mapping, key int) int {
	for _, trans := range currmap {
		ok, val := isIn(trans, key)
		if ok {
			return val
		}
	}
	return key
}

func isIn(detail mapping, key int) (bool, int) {
	if key >= detail.source && key < detail.source+detail.length {
		return true, detail.dest + key - detail.source
	}
	return false, -1
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
