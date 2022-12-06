package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tomschdev/advent-of-code/2022/go/day1/pkg/io"
)

var hardcodeResult = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}
var chooseCorrectPlay = map[string]map[string]string{
	"A": {
		"X": "Z",
		"Y": "X",
		"Z": "Y",
	},
	"B": {
		"X": "X",
		"Y": "Y",
		"Z": "Z",
	},
	"C": {
		"X": "Y",
		"Y": "Z",
		"Z": "X",
	},
}
var scorePerPlay = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}
var referee = map[string]map[string]int{
	"A": {
		"X": 3,
		"Y": 6,
		"Z": 0,
	},
	"B": {
		"X": 0,
		"Y": 3,
		"Z": 6,
	},
	"C": {
		"X": 6,
		"Y": 0,
		"Z": 3,
	},
}

func main() {
	contents := io.ReadInputBytes("")
	rounds := strings.Split(strings.Trim(string(contents), " "+"\n"), "\n")
	question1(rounds)
	question2(rounds)
}

func question1(rounds []string) {
	score := 0
	for _, r := range rounds {
		score += scoreRound(strings.Split(r, " "))
	}
	fmt.Println(score)
}

func question2(rounds []string) {
	score := 0
	for _, r := range rounds {
		score += hackRound(strings.Split(r, " "))
	}
	fmt.Println(score)
}

func scoreRound(plays []string) int {
	if len(plays) != 2 {
		fmt.Println("expected 2 plays, got: ", plays)
		os.Exit(1)
	}
	return referee[plays[0]][plays[1]] + scorePerPlay[plays[1]]
}

func hackRound(plays []string) int {
	if len(plays) != 2 {
		fmt.Println("expected 2 plays, got: ", plays)
		os.Exit(1)
	}
	return hardcodeResult[plays[1]] + scorePerPlay[chooseCorrectPlay[plays[0]][plays[1]]]
}
