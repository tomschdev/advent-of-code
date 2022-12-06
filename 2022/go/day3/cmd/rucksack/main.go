package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/tomschdev/advent-of-code/2022/go/day1/pkg/io"
)

func main() {
	contents := io.ReadInputBytes("")
	input := strings.Split(io.BytesToTrimmedString(contents, " "+"\n"), "\n")
	question1(input)
	question2(input)
}

func question2(input []string) {
	total := 0
	for i := 0; i < len(input)/3; i++ {
		r2 := findDuplicate(input[i*3], input[i*3+1])
		r3 := findDuplicate(string(r2), input[i*3+2])
		total += mapItemToPririoty(r3[0])
	}
	fmt.Println(total)
}

func question1(input []string) {
	total := 0
	for _, v := range input {
		r := findDuplicate(splitInputString(v))
		total += mapItemToPririoty(r[0])
	}
	fmt.Println(total)
}

func splitInputString(input string) (string, string) {
	if len(input)%2 != 0 {
		fmt.Println("input string is not of even length")
		os.Exit(1)
	}
	return input[:len(input)/2], input[len(input)/2:]
}

func findDuplicate(one, two string) []rune {
	runesOne, runesTwo := []rune(one), []rune(two)
	sort.Slice(runesOne, func(i, j int) bool {
		return runesOne[i] < runesOne[j]
	})
	sort.Slice(runesTwo, func(i, j int) bool {
		return runesTwo[i] < runesTwo[j]
	})
	var duplicates []rune
	for _, r1 := range runesOne {
		for _, r2 := range runesTwo {
			if r1 == r2 {
				duplicates = append(duplicates, r1)
			}
		}
	}
	return duplicates
}

func mapItemToPririoty(item rune) int {
	if int(item) <= 122 && int(item) >= 97 {
		return int(item) - 96
	} else if int(item) <= 90 && int(item) >= 65 {
		return int(item) - 38
	} else {
		return 0
	}
}
