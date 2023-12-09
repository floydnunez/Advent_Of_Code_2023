package main

import (
	"2023/util"
	"fmt"
	"strconv"
	"strings"
)

type data struct {
	original  []int
	sequences [][]int
}

func main() {
	fileLines := util.ReadFileIntoArray("day09/input.txt")
	var datas []data
	for _, line := range fileLines {
		numbers := strings.Split(line, " ")
		currdata := data{}
		for _, strnumber := range numbers {
			number, _ := strconv.Atoi(strnumber)
			currdata.original = append(currdata.original, number)
		}
		datas = append(datas, currdata)
	}
	//calc differences until all zeroes
	for i, currdata := range datas {
		currsequence := currdata.original
		currdata.sequences = append(currdata.sequences, currdata.original)
		for {
			differences, allZeros := calcDifferences(currsequence)
			currdata.sequences = append(currdata.sequences, differences)
			if allZeros {
				break
			}
			currsequence = differences
		}
		fmt.Println(currdata)
		datas[i] = currdata
	}
	//generate values
	fmt.Println("-------------------")
	totalP1 := 0
	for _, currdata := range datas {
		generatedVal := 0
		seqlen := len(currdata.sequences)
		for ii := 1; ii <= seqlen; ii++ {
			currdata.sequences[seqlen-ii] = append(currdata.sequences[seqlen-ii], generatedVal)
			if ii == seqlen {
				totalP1 += generatedVal
				break
			}
			nextseq := currdata.sequences[seqlen-ii-1]
			generatedVal += nextseq[len(nextseq)-1]
		}
		fmt.Println(currdata)
	}
	fmt.Println("Part 1: ", totalP1)
	//pre-generate values
	fmt.Println("-------------------")
	totalP2 := 0
	for _, currdata := range datas {
		generatedVal := 0
		seqlen := len(currdata.sequences)
		for ii := 1; ii <= seqlen; ii++ {
			currdata.sequences[seqlen-ii] = prepend(generatedVal, currdata.sequences[seqlen-ii])
			if ii == seqlen {
				totalP2 += generatedVal
				break
			}
			nextseq := currdata.sequences[seqlen-ii-1]
			generatedVal = nextseq[0] - generatedVal
		}
		fmt.Println(currdata)
	}
	fmt.Println("Part 2: ", totalP2)
}

func prepend(val int, ints []int) []int {
	return append([]int{val}, ints...)
}

func calcDifferences(sequence []int) ([]int, bool) {
	var result []int
	allZeros := true
	for i := 1; i < len(sequence); i++ {
		nextVal := sequence[i] - sequence[i-1]
		result = append(result, nextVal)
		if nextVal != 0 {
			allZeros = false
		}
	}
	return result, allZeros
}
