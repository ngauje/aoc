package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func indexesAll(str, substr string) []int {
	var resPos []int
	pos := strings.Index(str, substr)

	for pos != -1 {
		resPos = append(resPos, pos)
		pos = strings.Index(str[pos+len(substr):], substr)

		if pos != -1 {
			pos += len(substr)
			pos += resPos[len(resPos)-1]
		}
	}

	return resPos
}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo2")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	intRe := regexp.MustCompile(`(\d)`)

	result := 0

	stringNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	stringNumbersVal := make(map[string]int)
	stringNumbersVal["one"] = 1
	stringNumbersVal["two"] = 2
	stringNumbersVal["three"] = 3
	stringNumbersVal["four"] = 4
	stringNumbersVal["five"] = 5
	stringNumbersVal["six"] = 6
	stringNumbersVal["seven"] = 7
	stringNumbersVal["eight"] = 8
	stringNumbersVal["nine"] = 9

	for scanner.Scan() {

		line := scanner.Text()

		// allIntLine[pos] = val
		allIntLine := make(map[int]int)

		intMatches := intRe.FindAllStringIndex(line, -1)

		for _, valPos := range intMatches {
			val, _ := strconv.Atoi(line[valPos[0]:valPos[1]])
			allIntLine[valPos[0]] = val

		}

		for _, stringNumber := range stringNumbers {
			val := stringNumbersVal[stringNumber]
			// strings.Index find only the first match
			// indexesAll get all indexes for a match
			stringNumberpos := indexesAll(line, stringNumber)

			for _, pos := range stringNumberpos {
				allIntLine[pos] = val

				fmt.Printf("%v - %v\n", stringNumber, val)
			}
		}

		fmt.Printf("%+v\n", allIntLine)

		// hashmaps are not ordered
		posKeys := make([]int, 0, len(allIntLine))
		for p := range allIntLine {
			posKeys = append(posKeys, p)
		}

		sort.Ints(posKeys)

		val, _ := strconv.Atoi(strconv.Itoa(allIntLine[posKeys[0]]) + strconv.Itoa(allIntLine[posKeys[len(posKeys)-1]]))

		result = result + val
		fmt.Printf("%v\n", val)
	}

	fmt.Printf("Result %v\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
