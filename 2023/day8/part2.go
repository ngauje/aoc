package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type dir struct {
	L string
	R string
}

var mapDirection = map[string]dir{}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo3")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	directions := scanner.Text()
	scanner.Scan()

	mapRe := regexp.MustCompile(`([0-9A-Z]{3})`)

	curPos := []string{}
	totalSteps := []int{}

	for scanner.Scan() {
		mapLine := scanner.Text()
		mapDef := mapRe.FindAllString(mapLine, -1)

		if string(mapDef[0][2]) == "A" {
			curPos = append(curPos, mapDef[0])
		}

		mapDirection[mapDef[0]] = dir{
			L: mapDef[1],
			R: mapDef[2],
		}
	}

	steps := 0
	end := false

	for j := 0; j < len(curPos); j++ {
		steps = 0
		end = false

		for !end {
			for i := 0; i < len(directions); i++ {
				end = true

				switch string(directions[i]) {

				case "R":
					curPos[j] = string(mapDirection[curPos[j]].R)

					if string(curPos[j][2]) != "Z" {
						end = false
					}
				case "L":
					curPos[j] = string(mapDirection[curPos[j]].L)

					if string(curPos[j][2]) != "Z" {
						end = false
					}
				}
				steps++

				if end {
					break
				}
			}
		}
		fmt.Printf("Steps: %d\n", steps)
		totalSteps = append(totalSteps, steps)
	}

	fmt.Printf("LCM Steps: %d\n", LCM(totalSteps[0], totalSteps[1], totalSteps[2:]...))
}
