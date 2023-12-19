package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func followSeq(curSeq []int) int {
	end := true
	for i := 0; i < len(curSeq); i++ {
		if curSeq[i] != 0 {
			end = false
		}
	}

	if !end {
		var newSeq []int

		for i := 0; i < len(curSeq)-1; i++ {
			newSeq = append(newSeq, curSeq[i+1]-curSeq[i])
		}

		lastVal := followSeq(newSeq)

		return curSeq[len(curSeq)-1] + lastVal

	} else {
		return 0
	}
}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineInts := strings.Split(line, " ")

		var seq []int

		for i := 0; i < len(lineInts); i++ {
			val, _ := strconv.Atoi(lineInts[i])
			seq = append(seq, val)
		}

		res = res + followSeq(seq)

	}
	fmt.Printf("Result: %v\n", res)
}
