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

	var timeS, distS string

	for i := 0; i < len(timeMatches); i++ {
		timeS = timeS + timeMatches[i]
		distS = distS + distMatches[i]
	}

	time, _ := strconv.Atoi(timeS)
	distM, _ := strconv.Atoi(distS)

	waysToBeat := 1
	var currentWaysToBeat int

	// fmt.Printf("Time: %d\n", time)

	currentWaysToBeat = 0
	//
	for j := 0; j <= time; j++ {
		hold := j
		dist := (time - hold) * hold
		if dist > distM {
			currentWaysToBeat = currentWaysToBeat + 1
			// fmt.Printf("Race: %d - %d\n", hold, dist)
		}
	}
	waysToBeat = waysToBeat * currentWaysToBeat

	fmt.Printf("Result: %d\n", waysToBeat)
}
