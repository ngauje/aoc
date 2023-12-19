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

func findDirs(pipes [][]string, curPos coord) (coord, coord) {
	var top, bottom, left, right, cur string

	cur = pipes[curPos.y][curPos.x]

	if curPos.y-1 >= 0 {
		top = pipes[curPos.y-1][curPos.x]
	} else {
		top = "."
	}

	if curPos.y+1 < len(pipes) {
		bottom = pipes[curPos.y+1][curPos.x]
	} else {
		bottom = "."
	}

	if curPos.x-1 >= 0 {
		left = pipes[curPos.y][curPos.x-1]
	} else {
		left = "."
	}

	if curPos.x+1 < len(pipes[curPos.y]) {
		right = pipes[curPos.y][curPos.x+1]
	} else {
		right = "."
	}

	possibleCoords := []coord{}

	if ((cur != "-") && (cur != "7") && (cur != "F")) && ((top == "|") || (top == "F") || (top == "7") || (top == "S")) {
		possibleCoords = append(possibleCoords, coord{y: curPos.y - 1, x: curPos.x})
	}
	if ((cur != "-") && (cur != "L") && (cur != "J")) && ((bottom == "|") || (bottom == "L") || (bottom == "J") || (bottom == "S")) {
		possibleCoords = append(possibleCoords, coord{y: curPos.y + 1, x: curPos.x})
	}
	if ((cur != "|") && (cur != "L") && (cur != "F")) && ((left == "-") || (left == "F") || (left == "L") || (left == "S")) {
		possibleCoords = append(possibleCoords, coord{y: curPos.y, x: curPos.x - 1})
	}
	if ((cur != "|") && (cur != "7") && (cur != "J")) && ((right == "-") || (right == "7") || (right == "J") || (right == "S")) {
		possibleCoords = append(possibleCoords, coord{y: curPos.y, x: curPos.x + 1})
	}

	return possibleCoords[0], possibleCoords[1]
}

func isCorner(pipes [][]string, curPos coord) bool {
	curPipe := pipes[curPos.y][curPos.x]

	if (curPipe == "S") || (curPipe == "J") || (curPipe == "L") || (curPipe == "7") || (curPipe == "F") {
		return true
	}

	return false
}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo3")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pipes := [][]string{}

	scanner := bufio.NewScanner(file)

	y := 0
	start := coord{}
	for scanner.Scan() {
		line := scanner.Text()

		pipes = append(pipes, []string{})

		for x := 0; x < len(line); x++ {
			pipes[y] = append(pipes[y], string(line[x]))
			if string(line[x]) == "S" {
				start = coord{x: x, y: y}
			}
		}
		y++
	}

	step := 1
	end := false

	dir1 := coord{}
	dir2 := coord{}

	dir1, dir2 = findDirs(pipes, start)

	prevPos1 := start
	prevPos2 := start

	corners1 := []coord{start}
	corners2 := []coord{start}

	for !end {
		step++
		newDir11, newDir12 := findDirs(pipes, dir1)
		newDir21, newDir22 := findDirs(pipes, dir2)

		if newDir11 == prevPos1 {
			prevPos1 = dir1
			dir1 = newDir12
		} else {
			prevPos1 = dir1
			dir1 = newDir11
		}

		if newDir21 == prevPos2 {
			prevPos2 = dir2
			dir2 = newDir22
		} else {
			prevPos2 = dir2
			dir2 = newDir21
		}

		if isCorner(pipes, dir1) {
			corners1 = append(corners1, dir1)
		}

		if isCorner(pipes, dir2) {
			corners2 = append(corners2, dir2)
		}

		if dir1 == start {
			end = true
		}
	}

	// Shoelace formula
	sum := 0
	for i := 0; i < len(corners1)-1; i++ {
		sum = sum + ((corners1[i].x * corners1[i+1].y) - (corners1[i].y * corners1[i+1].x))
	}

	area := int(math.Abs(float64(sum / 2)))

	// Pick's theorem
	pointsEnclosedInLoop := ((area * -1) + (step / 2) - 1) * -1

	// fmt.Printf("step: %+v \n", step)
	// fmt.Printf("area: %+v \n", area)

	fmt.Printf("Points enclosed: %+v \n", pointsEnclosedInLoop)

}
