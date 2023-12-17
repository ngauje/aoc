package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type dir struct {
	L string
	R string
}

var mapDirection = map[string]dir{}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo")
	// file, err := os.Open("./input-demo2")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	directions := scanner.Text()
	scanner.Scan()

	mapRe := regexp.MustCompile(`([A-Z]{3})`)

	curPos := "AAA"

	for scanner.Scan() {
		mapLine := scanner.Text()
		mapDef := mapRe.FindAllString(mapLine, -1)

		mapDirection[mapDef[0]] = dir{
			L: mapDef[1],
			R: mapDef[2],
		}
	}

	steps := 0
	for curPos != "ZZZ" {
		for i := 0; i < len(directions); i++ {
			// fmt.Printf("Current position: %s, going %s\n", curPos, string(directions[i]))

			switch string(directions[i]) {

			case "R":
				curPos = string(mapDirection[curPos].R)
			case "L":
				curPos = string(mapDirection[curPos].L)
			}
			steps++

			if curPos == "ZZZ" {
				break
			}
		}
	}
	fmt.Printf("Steps: %d\n", steps)

}
