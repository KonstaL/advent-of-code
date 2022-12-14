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

	prioritySums := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {

		// Init rucksack comparments in bytes
		rucksackBytes := []byte(lines[i])
		compartmentDivider := len(rucksackBytes) / 2

		compA := rucksackBytes[:compartmentDivider]
		compB := rucksackBytes[compartmentDivider:]

		compABytes := make(map[byte]bool)

		// Add bytes to map to prevent duplicates
		for j := 0; j < len(compA); j++ {
			compABytes[compA[j]] = true
		}

		// Accumulate matches found to map, to prevent duplicates (ie. ABC XBB)
		accBytes := make(map[byte]bool)

		for j := 0; j < len(compB); j++ {
			if compABytes[compB[j]] {
				fmt.Println("Found a match!", compB[j])
				accBytes[compB[j]] = true
			}
		}

		// Sum found item types per row
		sum := 0
		for key := range accBytes {
			sum += byteToPriority(key)
		}
		prioritySums[i] = sum
	}

	summedPriorities := 0

	for i := 0; i < len(prioritySums); i++ {
		summedPriorities += prioritySums[i]
	}

	fmt.Println(summedPriorities)

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
