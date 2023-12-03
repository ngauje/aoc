package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type coord struct {
	x int
	y int
}

type part struct {
	// [[1,1], [2,1]]
	coords []coord
	value  int
	isPart bool
}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	partRe := regexp.MustCompile(`\d+`)
	symbolRe := regexp.MustCompile(`[^\d\.]`)

	y := 0

	partList := make(map[int][]*part)
	var symbolList []coord

	for scanner.Scan() {

		// fmt.Println("Line", y)
		currLine := scanner.Text()
		partMatches := partRe.FindAllStringSubmatchIndex(currLine, -1)
		if partMatches != nil {

			for _, p := range partMatches {

				// fmt.Printf("%+v\n", p)
				// fmt.Printf("Value: %+v\n", currLine[p[0]:p[1]])

				partVal, _ := strconv.Atoi(currLine[p[0]:p[1]])

				currPart := part{
					value: partVal,
				}

				for x := p[0]; x < p[1]; x++ {
					currPart.coords = append(currPart.coords, coord{x: x, y: y})
				}

				partList[y] = append(partList[y], &currPart)

				// fmt.Printf("%+v\n", currPart)

			}
		}

		symbolMatches := symbolRe.FindAllStringSubmatchIndex(currLine, -1)
		if symbolMatches != nil {
			for _, s := range symbolMatches {
				symbolList = append(symbolList, coord{y: y, x: s[0]})
			}
		}

		y = y + 1
	}

	// Loop on all symbol to find adjacent parts
	for _, s := range symbolList {
		for y := s.y - 1; y <= s.y+1; y++ {
			// for x := s.x - 1; x <= s.x+1; x++ {
			for _, part := range partList[y] {
				for _, partCoord := range part.coords {
					if partCoord.x == s.x || partCoord.x == s.x-1 || partCoord.x == s.x+1 {
						part.isPart = true
					}
				}
			}
			// }
		}
	}

	partSum := 0
	for _, y := range partList {
		for _, part := range y {
			if part.isPart {
				partSum = partSum + part.value
			}
		}
	}

	fmt.Printf("Total: %+v\n", partSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
