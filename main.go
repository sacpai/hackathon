package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Printf("Start time: %s\n", startTime)
	data, err := os.ReadFile("data/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	//sorting lines
	sort.Strings(lines)
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
	fmt.Printf("{")
	looper := 0
	//calculate values
	for city, temps := range dataStore {
		//fmt.Printf("%s: ", key)
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

	fmt.Printf("Time taken: %s\n", time.Since(startTime))

	//fmt.Println(string(data))
}
