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

func parseRange(rawVal string) [3]int {
	var values [3]int
	rangeVal := strings.Split(rawVal, " ")

	values[0], _ = strconv.Atoi(rangeVal[0])
	values[1], _ = strconv.Atoi(rangeVal[1])
	values[2], _ = strconv.Atoi(rangeVal[2])
	return values
}

func main() {

	file, err := os.Open("./input")
	// file, err := os.Open("./input-demo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var currentBloc string

	seeds := make([]int, 0)

	// blocs regexp
	seedsRe := regexp.MustCompile(`^seeds: .*`)
	seedSoilRe := regexp.MustCompile(`^seed-to-soil map:$`)
	soilFertilizerRe := regexp.MustCompile(`^soil-to-fertilizer map:$`)
	fertilizerWaterRe := regexp.MustCompile(`^fertilizer-to-water map:$`)
	waterLightRe := regexp.MustCompile(`^water-to-light map:$`)
	lightTemperatureRe := regexp.MustCompile(`^light-to-temperature map:$`)
	temperatureHumidityRe := regexp.MustCompile(`^temperature-to-humidity map:$`)
	humidityLocationRe := regexp.MustCompile(`^humidity-to-location map:$`)

	// maps
	seedMaps := make(map[int][][3]int)

	currentMap := 0

	for scanner.Scan() {

		if scanner.Text() != "" {
			newBloc := false

			if seedsRe.MatchString(scanner.Text()) {
				// currentBloc = "seeds"
				newBloc = true
				seedsVal := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
				for i := 0; i < len(seedsVal); i++ {
					v, _ := strconv.Atoi(seedsVal[i])
					seeds = append(seeds, v)
				}
			}

			if seedSoilRe.MatchString(scanner.Text()) {
				// currentBloc = "seedSoil"
				newBloc = true
				currentMap = 0
			}

			if soilFertilizerRe.MatchString(scanner.Text()) {
				// currentBloc = "soilFertilizer"
				newBloc = true
				currentMap = 1
			}

			if fertilizerWaterRe.MatchString(scanner.Text()) {
				// currentBloc = "fertilizerWater"
				newBloc = true
				currentMap = 2
			}

			if waterLightRe.MatchString(scanner.Text()) {
				// currentBloc = "waterLight"
				newBloc = true
				currentMap = 3
			}

			if lightTemperatureRe.MatchString(scanner.Text()) {
				// currentBloc = "lightTemperature"
				newBloc = true
				currentMap = 4
			}

			if temperatureHumidityRe.MatchString(scanner.Text()) {
				// currentBloc = "temperatureHumidity"
				newBloc = true
				currentMap = 5
			}

			if humidityLocationRe.MatchString(scanner.Text()) {
				// currentBloc = "humidityLocation"
				newBloc = true
				currentMap = 6
			}

			if !newBloc {
				ranges := parseRange(scanner.Text())

				seedMaps[currentMap] = append(seedMaps[currentMap], ranges)

				// fmt.Printf("%+v\n", seedMaps)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	var lowestLocation int

	for i := 0; i < len(seeds); i++ {

		in := seeds[i]
		// fmt.Printf("IN: %d\n", in)
		var out int

		for j := 0; j < 7; j++ {

			outed := false
			for _, rangeMap := range seedMaps[j] {
				if !outed {
					if in >= rangeMap[1] && in < rangeMap[1]+rangeMap[2] {
						out = in - rangeMap[1] + rangeMap[0]
						outed = true
					}
				}
			}
			if !outed {
				out = in
			}

			// fmt.Printf("%d -> %d\n", in, out)

			in = out

		}
		// fmt.Printf("OUT: %d\n", in)

		if lowestLocation == 0 {
			lowestLocation = in
		} else {
			if in < lowestLocation {
				lowestLocation = in
			}
		}
	}

	fmt.Printf("lowestLocation: %d\n", lowestLocation)
}
