package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	totalScore := make(map[int]int)
	totalSum := 0

	cardNumRe := regexp.MustCompile(`(\d+)`)

	for scanner.Scan() {

		// card := make(map[int][]int)

		numbers := strings.Split(strings.Split(scanner.Text(), ":")[1], "|")

		numCardMatch := cardNumRe.FindStringSubmatch(strings.Split(scanner.Text(), ":")[0])
		numCard, _ := strconv.Atoi(numCardMatch[1])

		if _, f := totalScore[numCard]; f {
			totalScore[numCard] = totalScore[numCard] + 1
		} else {
			totalScore[numCard] = 1
		}

		myNumbers := strings.Split(numbers[1], " ")
		myNumbers = removeEmptyStrings(myNumbers)

		winnings := strings.Split(numbers[0], " ")
		winnings = removeEmptyStrings(winnings)

		score := 0

		for i := 0; i < len(myNumbers); i++ {

			for j := 0; j < len(winnings); j++ {
				if myNumbers[i] == winnings[j] {

					score = score + 1
				}

			}

		}

		for j := 0; j < totalScore[numCard]; j++ {
			for i := numCard + 1; i < score+numCard+1; i++ {

				if _, f := totalScore[i]; f {
					totalScore[i] = totalScore[i] + 1
				} else {
					totalScore[i] = 1
				}

			}
		}

		totalSum = totalSum + totalScore[numCard]

		// fmt.Printf("%+v\n", numbers[0])
	}

	fmt.Printf("Total: %+v\n", totalSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
