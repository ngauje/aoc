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

	intRe := regexp.MustCompile(`(\d)`)

	result := 0

	for scanner.Scan() {

		line := scanner.Text()

		intMatches := intRe.FindAllString(line, -1)

		if intMatches != nil {

			val, _ := strconv.Atoi(intMatches[0] + intMatches[len(intMatches)-1])

			// fmt.Printf("%+v\n", intMatches[0]+intMatches[len(intMatches)-1])

			result = result + val
		}
	}

	fmt.Printf("Result %v\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
