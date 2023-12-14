package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeLine := scanner.Text()

	scanner.Scan()
	distLine := scanner.Text()

	valRe := regexp.MustCompile(`(\d+)`)

	timeMatches := valRe.FindAllString(timeLine, -1)
	distMatches := valRe.FindAllString(distLine, -1)

	var times []int
	var dists []int

	for i := 0; i < len(timeMatches); i++ {
		val, _ := strconv.Atoi(timeMatches[i])
		times = append(times, val)

		val, _ = strconv.Atoi(distMatches[i])
		dists = append(dists, val)
	}

	waysToBeat := 1
	var currentWaysToBeat int
	for i := 0; i < len(timeMatches); i++ {
		currentWaysToBeat = 0
		// fmt.Printf("Time: %d\n", times[i])
		for j := i; j <= times[i]; j++ {
			hold := j
			dist := (times[i] - hold) * hold
			if dist > dists[i] {
				currentWaysToBeat = currentWaysToBeat + 1
				fmt.Printf("Race: %d - %d\n", hold, dist)
			}

		}

		waysToBeat = waysToBeat * currentWaysToBeat
	}
	fmt.Printf("Result: %d\n", waysToBeat)
}
