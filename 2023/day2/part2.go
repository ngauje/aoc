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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type set struct {
	blue  int
	red   int
	green int
}

type game struct {
	sets []set
	id   int
}

var games = make([]game, 0)

var sumIds = 0
var sumPower = 0

// array = append(array, 1)

func main() {

	// file, err := os.Open("./input-demo")
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	setBlueRe := regexp.MustCompile(`[^\d]*(\d+) blue.*`)
	setRedRe := regexp.MustCompile(`[^\d](\d+) red.*`)
	setGreenRe := regexp.MustCompile(`[^\d](\d+) green.*`)

	for scanner.Scan() {

		currGame := new(game)

		maxReds := 0
		maxGreens := 0
		maxBlues := 0

		gameInfo := strings.Split(scanner.Text(), ":")
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		currGame.id, _ = strconv.Atoi(strings.Split(gameInfo[0], " ")[1])

		gameSets := strings.Split(gameInfo[1], ";")

		for _, s := range gameSets {
			currSet := set{
				blue:  0,
				red:   0,
				green: 0,
			}

			// fmt.Println(s)
			// s= "1 red, 2 green, 6 blue"

			blueMatches := setBlueRe.FindStringSubmatch(s)

			if blueMatches != nil {
				currSet.blue, _ = strconv.Atoi(string(blueMatches[1]))
				if currSet.blue > maxBlues {
					maxBlues = currSet.blue
				}
			}

			redMatches := setRedRe.FindStringSubmatch(s)

			if redMatches != nil {
				currSet.red, _ = strconv.Atoi(string(redMatches[1]))
				if currSet.red > maxReds {
					maxReds = currSet.red
				}
			}

			greenMatches := setGreenRe.FindStringSubmatch(s)

			if greenMatches != nil {
				currSet.green, _ = strconv.Atoi(string(greenMatches[1]))
				if currSet.green > maxGreens {
					maxGreens = currSet.green
				}
			}

			currGame.sets = append(currGame.sets, currSet)

			// fmt.Printf("%+v\n", currSet)
		}

		power := maxReds * maxGreens * maxBlues

		sumPower = sumPower + power

		// fmt.Printf("%+v\n", currGame)
		// fmt.Printf("%+v\n", power)
	}

	fmt.Printf("Total power: %d\n", sumPower)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
