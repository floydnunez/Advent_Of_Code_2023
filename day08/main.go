package main

import (
	"2023/util"
	"fmt"
	"strings"
)

type node struct {
	name string
	l    string
	r    string
}

func main() {
	fileLines := util.ReadFileIntoArray("day08/input.txt")
	partB := true
	nodes := make(map[string]node)
	instructions := ""
	for i, line := range fileLines {
		if i == 0 {
			instructions = line
			continue
		} else if i == 1 {
			continue
		}
		name, left, right := parseLine(line)
		nodes[name] = node{name: name, l: left, r: right}
	}
	fmt.Println("instructions:", instructions)
	fmt.Println("nodes:", nodes)

	if !partB {
		currnode := nodes["AAA"]
		count := 0
		var isZ bool
		fmt.Println("primer nodo:", currnode)
		for i := 0; ; i++ {
			i = i % len(instructions)
			count++
			currnode, isZ = step(instructions, i, nodes, currnode, partB)
			if isZ {
				break
			}
		}
		fmt.Println("Part 1", count, currnode) //22357
	} else {
		var ini []node
		for key := range nodes {
			if key[2:3] == "A" {
				ini = append(ini, nodes[key])
			}
		}
		fmt.Println("ini:", ini)
		var results []int

		for kk := 0; kk < len(ini); kk++ {
			currnode := ini[kk]
			count := 0
			var isZ bool
			fmt.Println("primer nodo:", currnode)
			for i := 0; ; i++ {
				i = i % len(instructions)
				count++
				currnode, isZ = step(instructions, i, nodes, currnode, partB)
				if isZ {
					break
				}
			}
			results = append(results, count)
			ini[kk] = currnode
		}

		fmt.Println("count:", results)
		fmt.Println("Part 2:", LCM(results)) //10371555451871
	}
}

func step(instructions string, i int, nodes map[string]node, currnode node, partB bool) (node, bool) {
	inst := instructions[i]
	var newnode node
	if inst == 'L' {
		newnode = nodes[currnode.l]
	} else {
		newnode = nodes[currnode.r]
	}
	//fmt.Println(currnode, string(inst), " => ", newnode)
	currnode = newnode
	isZ := false
	if !partB && currnode.name == "ZZZ" {
		isZ = true
	} else if partB {
		isZ = currnode.name[2:3] == "Z"
	}
	return currnode, isZ
}

func parseLine(line string) (string, string, string) {
	parts := strings.Split(line, " ")
	name := parts[0]
	left := parts[2][1:4]
	right := parts[3][:3]
	fmt.Println("parts:", parts, name, left, right)
	return name, left, right
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers []int) int {
	if len(integers) < 2 {
		return 0
	}
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])
	for i := 2; i < len(integers); i++ {
		result = result * integers[i] / GCD(result, integers[i])
	}
	return result
}
