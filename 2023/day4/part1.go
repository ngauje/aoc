package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0

	for scanner.Scan() {

		// card := make(map[int][]int)

		numbers := strings.Split(strings.Split(scanner.Text(), ":")[1], "|")

		// winnings := numbers[0]
		myNumbers := strings.Split(numbers[1], " ")
		myNumbers = removeEmptyStrings(myNumbers)

		winnings := strings.Split(numbers[0], " ")
		winnings = removeEmptyStrings(winnings)

		first := true
		score := 0

		for i := 0; i < len(myNumbers); i++ {
			// fmt.Printf("Number: %+v\n", myNumbers[i])

			for j := 0; j < len(winnings); j++ {
				if myNumbers[i] == winnings[j] {
					// fmt.Printf("Match: %+v\n", winnings[j])

					if first {
						score = 1
						first = false
					} else {
						score = score * 2
					}

				}

			}
		}
		// fmt.Printf("%+v\n", score)
		totalScore = totalScore + score

		// fmt.Printf("%+v\n", numbers[0])
	}

	fmt.Printf("Total: %+v\n", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
