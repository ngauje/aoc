package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards        []byte
	cardsInValue []int
	bid          int
	handType     int
}

var cardsValue = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardList = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

func getHandType(cards []byte) int {
	pairs := 0
	brelan := 0
	four := 0
	five := 0
	for i := 0; i < len(cardList); i++ {
		total := strings.Count(string(cards), cardList[i])

		switch total {
		case 2:
			pairs = pairs + 1
		case 3:
			brelan = brelan + 1
		case 4:
			four = four + 1
		case 5:
			five = five + 1
		}
	}

	if five == 1 {
		return 6 // Five of a kind: AAAAA
	} else if four == 1 {
		return 5 // Four of a kind: AA8AA
	} else if brelan == 1 && pairs == 1 {
		return 4 // Full house: 23332
	} else if brelan == 1 && pairs == 0 {
		return 3 // Three of a kind: TTT98
	} else if pairs == 2 {
		return 2 // Two pair: 23432
	} else if pairs == 1 {
		return 1 // One pair: A23A4
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

	var hands = []hand{}

	for scanner.Scan() {

		line := scanner.Text()
		splittedLine := strings.Split(line, " ")

		cards := splittedLine[0]
		bid, _ := strconv.Atoi(splittedLine[1])

		var cardsVal []int

		for i := 0; i < 5; i++ {

			cardsVal = append(cardsVal, cardsValue[cards[i]])
		}

		newHand := hand{
			cards:        []byte(cards),
			cardsInValue: cardsVal,
			bid:          bid,
			handType:     getHandType([]byte(cards)),
		}

		hands = append(hands, newHand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		} else {
			k := 0
			for k = 0; k < 5; k++ {
				if hands[i].cardsInValue[k] == hands[j].cardsInValue[k] {
					continue
				} else {
					break
				}
			}
			return hands[i].cardsInValue[k] < hands[j].cardsInValue[k]
		}
	})

	total := 0
	for i := 0; i < len(hands); i++ {
		total = total + (hands[i].bid * (i + 1))
	}

	fmt.Printf("Total: %d\n", total)
}
