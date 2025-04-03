package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type measurements struct {
	City  string
	Temps []float64
}

func main() {
	data, err := os.ReadFile("data/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	//store the values in map
	dataStore := make(map[string][]float64)
	for _, line := range lines {
		lineSplit := strings.Split(line, ";")
		if len(lineSplit) == 2 {
			city := lineSplit[0]
			temp, err := strconv.ParseFloat(lineSplit[1], 64)
			if err == nil {
				dataStore[city] = append(dataStore[city], temp)
			}
			//fmt.Println(string(city) + strconv.FormatFloat(temp, 'f', -1, 64))
		}
	}
	//Sort Cities
	var sortedCities []string
	for city := range dataStore {
		sortedCities = append(sortedCities, city)
	}
	sort.Strings(sortedCities)

	fmt.Printf("{")
	//calculate values
	for looper, city := range sortedCities {
		//fmt.Printf("%s: ", key)
		temps := dataStore[city]
		var minVal, maxVal, totalVal float64
		for _, val := range temps {
			if val < minVal {
				minVal = val
			}
			if val > maxVal {
				maxVal = val
			}
			totalVal += val
			//fmt.Printf("%s, %.2f,%.2f \n", key, minVal, maxVal)
		}
		meanVal := totalVal / float64(len(temps))
		fmt.Printf("%s=%.1f/%.1f/%.1f", city, minVal, maxVal, math.Round(meanVal*10)/10)

		if looper < len(dataStore)-1 {
			fmt.Printf(",")
		}
		looper++
	}
	fmt.Printf("}\n")
	//fmt.Println(string(data))
}
