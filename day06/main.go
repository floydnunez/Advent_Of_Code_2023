package main

import "fmt"

func main() {
	real := 3
	var times []int
	times = []int{7, 15, 30}
	var distances []int
	distances = []int{9, 40, 200}
	if real == 1 {
		times = []int{49, 87, 78, 95}
		distances = []int{356, 1378, 1502, 1882}
	}
	if real == 2 {
		times = []int{71530}
		distances = []int{940200}
	}
	if real == 3 {
		times = []int{49877895}
		distances = []int{356137815021882}
	}
	fmt.Println("real?", real, times, distances)
	var results []int
	for i, time := range times {
		maxDistance := distances[i]
		won := 0
		for kk := 1; kk < time; kk++ {
			traveled := calcDistanceTraveled(kk, time)
			if traveled > maxDistance {
				won++
			}
		}
		results = append(results, won)
		fmt.Println("----------------------------------------------")
	}
	fmt.Println(results)
	total := 1
	for _, result := range results {
		total *= result
	}
	fmt.Println("Part 1:", total)
}

func calcDistanceTraveled(pressed int, max int) int {
	speed := pressed
	timeTraveled := max - pressed
	total := speed * timeTraveled
	return total
}
