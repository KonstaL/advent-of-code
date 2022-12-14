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

	fullyContainedPairs := 0

	for _, line := range lines {
		ranges := strings.Split(line, ",")
		range1Arr := strings.Split(ranges[0], "-")
		range2Arr := strings.Split(ranges[1], "-")

		isFullyContained := aContainsB(range1Arr, range2Arr) || aContainsB(range2Arr, range1Arr)

		if isFullyContained {
			fullyContainedPairs++
		}
	}

	fmt.Println(fullyContainedPairs)
}

func aContainsB(a []string, b []string) bool {
	aBottom, _ := strconv.Atoi(a[0])
	aTop, _ := strconv.Atoi(a[1])
	bBottom, _ := strconv.Atoi(b[0])
	bTop, _ := strconv.Atoi(b[1])

	if aBottom <= bBottom && aTop >= bTop {
		return true
	}

	return false
}
