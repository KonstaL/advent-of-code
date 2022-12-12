package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}

	str := string(data)
	strArr := strings.Split(str, "\n")

	calList := make([]int, 0)
	cals := 0
	for i := 0; i < len(strArr); i++ {
		element := strArr[i]

		if element == "" {
			calList = append(calList, cals)
			cals = 0
			continue
		}

		val, err := strconv.Atoi(strArr[i])
		if err != nil {
			fmt.Printf("Could not convert to int due to this %s error \n", err)
		}
		cals += val
	}

	sort.Ints(calList)

	top3Acc := 0
	for i := 0; i < 3; i++ {
		top3Acc += calList[len(calList)-1-i]
	}

	fmt.Println(top3Acc)

}
