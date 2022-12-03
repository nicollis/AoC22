package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1Stream(fileLoc string) {
	//println("Advent of Code 2022 Day 1!")

	//println("Part 1")

	max_cals := 0
	current_cals := 0
	file, err := os.Open(fileLoc)

	if err != nil {
		fmt.Printf("Cound not open file: %s\n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		number, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			if max_cals < current_cals {
				max_cals = current_cals
			}
			current_cals = 0
		} else {
			current_cals += number
		}
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close file: %s\n", err)
	}

	//fmt.Printf("Max Calories: %d\n", max_cals)
}

func Day1Loaded(fileLoc string) {
	//println("Advent of Code 2022 Day 1!")
	//
	//println("Part 1")

	max_cals := 0
	current_cals := 0
	file, err := os.ReadFile(fileLoc)

	if err != nil {
		fmt.Printf("Cound not open file: %s\n", err)
	}

	fileContent := strings.Split(string(file), "\n")

	for _, text := range fileContent {
		number, err := strconv.Atoi(text)
		if err != nil {
			if max_cals < current_cals {
				max_cals = current_cals
			}
			current_cals = 0
		} else {
			current_cals += number
		}
	}

	//fmt.Printf("Max Calories: %d\n", max_cals)
}

func Day1Part2(fileLoc string) {
	println("Advent of Code 2022 Day 1!")

	println("Part 2")

	var cals []int
	currentCals := 0
	file, err := os.ReadFile(fileLoc)

	if err != nil {
		fmt.Printf("Cound not open file: %s\n", err)
	}

	fileContent := strings.Split(string(file), "\n")

	for _, text := range fileContent {
		number, err := strconv.Atoi(text)
		if err != nil {
			cals = append(cals, currentCals)
			currentCals = 0
		} else {
			currentCals += number
		}
	}

	sort.Slice(cals, func(i, j int) bool {
		return cals[i] > cals[j]
	})

	fmt.Printf("Top 3 Calories: %d, %d, %d\n", cals[0], cals[1], cals[2])

	for _, cal := range cals[:3] {
		currentCals += cal
	}

	fmt.Printf("Cal Sum: %d\n", currentCals)
}
