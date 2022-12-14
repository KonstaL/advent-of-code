package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}

	strFile := string(data)
	lines := strings.Split(strFile, "\n")

	groups := make([][]string, len(lines)/3)

	groupInsertIndex := 0
	for i := 0; i < len(lines); i++ {
		groups[groupInsertIndex] = append(groups[groupInsertIndex], lines[i])

		if i%3 == 2 {
			groupInsertIndex++
		}
	}

	prioritySum := 0

	for i := 0; i < len(groups); i++ {
		groupBadge := findGroupBadge(groups[i])
		prioritySum += byteToPriority(groupBadge)
	}

	fmt.Println(prioritySum)
}

func findGroupBadge(group []string) byte {
	charMap := make(map[byte]int)

	for i := 0; i < len(group); i++ {
		rucksack := group[i]
		rucksackMap := make(map[byte]bool)
		for j := 0; j < len(rucksack); j++ {
			rucksackMap[byte(rucksack[j])] = true
		}

		for key := range rucksackMap {
			charMap[key]++
		}
	}

	commonBadge := byte(0)

	for key, val := range charMap {
		if val == len(group) {
			commonBadge = key
		}
	}

	return commonBadge
}

// We'll use this to convert the byte to a priority. Assumes input of A-z
func byteToPriority(input byte) int {
	// If capital letter
	if input < 97 {
		// 65 is the ascii value of A, so well subtract that from the input and add 26 to get the correct priority (a = 1, b = 2, ... z = 26, A = 27, B = 28, ... Z = 52)
		return int(input) - 64 + 26
	}
	return int(input) - 96
}
