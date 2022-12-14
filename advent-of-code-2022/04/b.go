package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}

	strFile := string(data)
	lines := strings.Split(strFile, "\n")

	overlappingPairs := 0

	for _, line := range lines {
		ranges := strings.Split(line, ",")
		range1Arr := strings.Split(ranges[0], "-")
		range2Arr := strings.Split(ranges[1], "-")

		areRangesOverlapping := doesRangesOverlap(range1Arr, range2Arr) || doesRangesOverlap(range2Arr, range1Arr)

		if areRangesOverlapping {
			overlappingPairs++
		}
	}

	fmt.Println(overlappingPairs)
}

func doesRangesOverlap(a []string, b []string) bool {
	aBottom, _ := strconv.Atoi(a[0])
	aTop, _ := strconv.Atoi(a[1])
	bBottom, _ := strconv.Atoi(b[0])
	bTop, _ := strconv.Atoi(b[1])

	if bBottom >= aBottom && bBottom <= aTop {
		return true
	}

	if bTop >= aBottom && bTop <= aTop {
		return true
	}

	return false
}
