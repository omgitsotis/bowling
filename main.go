package main

import (
	// "bufio"
	// "os"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	score, _ := readScore("X X X X X X X X X XXX")
	fmt.Println(score)
}

func readScore(input string) (int, error) {
	scores := strings.Split(input, " ")
	if len(scores) != 10 {
		return 0, errors.New("incorrect input card")
	}

	gameScore := make([]int, 10)
	totalScore := 0
	// wasStrike := false
	// wasSpare := false

	for i, score := range scores {
		// scoreValue := 0
		switch score {
		case "X":
			gameScore[i] = 10
		case "-":
			gameScore[i] = 0
		case "/":
			gameScore[i] = 10
		default:
			gameScore[i], _ = strconv.Atoi(score)
		}
	}

	for j, frame := range gameScore {
		if j <
	}

	return totalScore, nil
}
