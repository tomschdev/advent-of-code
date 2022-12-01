package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/tomschdev/advent-of-code/2022/go/day1/pkg/io"
)

func main() {
	contents := io.ReadInputBytes("")
	input := string(contents)
	calories := countCaloriesPerElf(input)
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	fmt.Println(calories[0] + calories[1] + calories[2])
}

func countCaloriesPerElf(input string) []int64 {
	input = strings.Trim(input, " "+"\n")

	calorieSlice := make([]int64, 0)
	if input == "" {
		return calorieSlice
	}
	grouping := strings.Split(input, "\n\n")

	for _, collection := range grouping {
		calorieEntries := strings.Split(collection, "\n")
		calorieCount := 0
		for _, c := range calorieEntries {
			cInt, err := strconv.ParseInt(strings.Trim(c, " "), 10, 64)
			if err != nil {
				fmt.Println("could not parse string to int64: ", c)
				os.Exit(1)
			}
			calorieCount += int(cInt)
		}
		calorieSlice = append(calorieSlice, int64(calorieCount))
	}
	return calorieSlice
}
