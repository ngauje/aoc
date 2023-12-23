package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

var toExpandY []int
var toExpandX []int

var expandIncrement = 1000000 - 1

type coord struct {
	x int
	y int
}

func galaxiesDistance(gal1 coord, gal2 coord) int {

	g1 := gal1
	g2 := gal2

	if gal1.x < gal2.x {
		for i := gal1.x; i < gal2.x; i++ {
			if slices.Contains(toExpandX, i) {
				g2.x = g2.x + expandIncrement
			}
		}
	}
	if gal1.x > gal2.x {
		for i := gal2.x; i < gal1.x; i++ {
			if slices.Contains(toExpandX, i) {
				g1.x = g1.x + expandIncrement
			}
		}
	}

	if gal1.y < gal2.y {
		for i := gal1.y; i < gal2.y; i++ {
			if slices.Contains(toExpandY, i) {
				g2.y = g2.y + expandIncrement
			}
		}
	}
	if gal1.y > gal2.y {
		for i := gal2.y; i < gal1.y; i++ {
			if slices.Contains(toExpandY, i) {
				g1.y = g1.y + expandIncrement
			}
		}
	}

	dist := coord{
		x: g2.x - g1.x,
		y: g2.y - g1.y,
	}

	return int(math.Abs(float64(dist.x)) + math.Abs(float64(dist.y)))
}

func allGalaxiesDistances(galaxies []coord) []int {
	distances := []int{}

	for i := 1; i < len(galaxies); i++ {
		distances = append(distances, galaxiesDistance(galaxies[0], galaxies[i]))
	}

	if len(galaxies) > 2 {
		distances = append(distances, allGalaxiesDistances(galaxies[1:])...)
	}

	return distances
}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	space := [][]string{}

	scanner := bufio.NewScanner(file)

	y := 0
	sizeY := 0
	sizeX := 0

	galaxies := []coord{}

	for scanner.Scan() {
		line := scanner.Text()

		space = append(space, []string{})

		emptyLine := true

		if sizeX == 0 {
			sizeX = len(line)
		}

		for x := 0; x < len(line); x++ {
			space[y] = append(space[y], string(line[x]))
			if line[x] != '.' {
				galaxies = append(galaxies, coord{x: x, y: y})
				emptyLine = false
			}
		}

		if emptyLine {
			toExpandY = append(toExpandY, y)
		}

		y++
	}

	sizeY = y

	for x := 0; x < sizeX; x++ {

		emptyColumn := true
		for y = 0; y < sizeY; y++ {
			if space[y][x] != "." {
				emptyColumn = false
			}
		}

		if emptyColumn {
			toExpandX = append(toExpandX, x)
		}
	}

	for i := len(toExpandX) - 1; i >= 0; i-- {

		for y = 0; y < sizeY; y++ {
			space[y] = append(space[y][:toExpandX[i]+1], space[y][toExpandX[i]:]...)
			space[y][toExpandX[i]] = "."
		}

		sizeX++

	}

	for i := len(toExpandY) - 1; i >= 0; i-- {

		space = append(space[:toExpandY[i]+1], space[toExpandY[i]:]...)
		for x := 0; x < sizeX; x++ {
			space[toExpandY[i]][x] = "."
		}

		sizeY++

	}

	// for y = 0; y < sizeY; y++ {
	// 	for x := 0; x < sizeX; x++ {

	// 		fmt.Printf("%s", space[y][x])
	// 	}
	// 	fmt.Printf("\n")
	// }

	var distances []int
	total := 0

	distances = allGalaxiesDistances(galaxies)
	for i := 0; i < len(distances); i++ {
		total = total + distances[i]
	}

	fmt.Printf("Total distances: %d\n", total)
}
