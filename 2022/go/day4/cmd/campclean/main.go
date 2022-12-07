package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tomschdev/advent-of-code/2022/go/day1/pkg/io"
)

func main() {
	contents := io.ReadInputBytes("")
	input := strings.Split(io.BytesToTrimmedString(contents, " "+"\n"), "\n")
	question1(input)
	question2(input)
}

func question1(input []string) {
	fullyContainedAssignmentCount := 0
	for _, entry := range input {
		fullyContainedAssignmentCount += checkFullyContainedAssignment(splitEntryToInts(entry))
	}
	fmt.Println(fullyContainedAssignmentCount)
}
func question2(input []string) {
	overlappingAssignmentCount := 0
	for _, entry := range input {
		overlappingAssignmentCount += checkOverlappingAssignment(splitEntryToInts(entry))
	}
	fmt.Println(overlappingAssignmentCount)
}

func splitEntryToInts(entry string) map[int]int { // must be more concise way to do this using mapping functions/list comprehensions
	elfSections := strings.Split(entry, ",")
	numericSections := make(map[int]int, 0)
	count := 0
	for _, e := range elfSections {
		strSects := strings.Split(e, "-")
		for _, s := range strSects {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("string to interger conversion failed")
				os.Exit(1)
			}
			numericSections[count] = n
			count++
		}
	}
	return numericSections
}

func checkFullyContainedAssignment(sectMap map[int]int) int {
	if sectMap[0] <= sectMap[2] && sectMap[1] >= sectMap[3] || sectMap[2] <= sectMap[0] && sectMap[3] >= sectMap[1] {
		return 1
	}
	return 0
}
func checkOverlappingAssignment(sectMap map[int]int) int {
	if sectMap[1] >= sectMap[2] && sectMap[0] <= sectMap[2] || sectMap[0] <= sectMap[3] && sectMap[1] >= sectMap[3] || checkFullyContainedAssignment(sectMap) == 1 {
		return 1
	}
	return 0
}
