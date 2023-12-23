package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type coord struct {
	x int
	y int
}

func galaxiesDistance(g1 coord, g2 coord) int {
	dist := coord{
		x: g2.x - g1.x,
		y: g2.y - g1.y,
	}

	steps := 0

	x := 0.0
	y := 0.0
	for x < math.Abs(float64(dist.x)) || y < math.Abs(float64(dist.y)) {
		if ((x + 0.5) / math.Abs(float64(dist.x))) < ((y + 0.5) / math.Abs(float64(dist.y))) {
			steps = steps + 1
			x = x + 1
		} else {
			steps = steps + 1
			y = y + 1
		}
	}

	return steps
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

	toExpandY := []int{}
	toExpandX := []int{}

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

	galaxies := []coord{}
	for y = 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			if space[y][x] == "#" {
				galaxies = append(galaxies, coord{x: x, y: y})
			}
			// fmt.Printf("%s", space[y][x])
		}
		// fmt.Printf("\n")
	}

	// fmt.Printf("%+v\n", galaxies)

	var distances []int
	total := 0

	distances = allGalaxiesDistances(galaxies)
	for i := 0; i < len(distances); i++ {
		total = total + distances[i]
	}

	fmt.Printf("Total distances: %d\n", total)

}
