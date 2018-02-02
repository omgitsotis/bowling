package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		fmt.Print("Would you like to start the game? ")
		line := readLine()

		switch strings.ToLower(line) {
		case "yes", "y":
			startGame()
		case "no", "n":
			os.Exit(0)
		}
	}

}

func startGame() {
	rand.Seed(time.Now().Unix())

	var playerScoreCard string
	var cpuScoreCard string
	var playerScore int
	var cpuScore int

	frameCount := 1
	for {
		playerScoreCard = playFrame(playerScoreCard)
		fmt.Printf("player score so far [%s]\n", playerScoreCard)
		cpuScoreCard = playFrame(cpuScoreCard)
		fmt.Printf("cpu score card so far[%s]\n", cpuScoreCard)
		fmt.Print("Continue? ")
		line := readLine()
		switch strings.ToLower(line) {
		case "yes":
			frameCount++
		default:
			os.Exit(0)
		}

		if frameCount > 9 {
			break
		}
	}

	var finalScore string
	if playerScore > cpuScore {
		finalScore = fmt.Sprintf("player wins [%d] to [%d]\n", playerScore, cpuScore)
	} else {
		finalScore = fmt.Sprintf("cpu wins [%d] to [%d]\n", cpuScore, playerScore)
	}

	fmt.Printf(finalScore)
}

func playGame() (string, int) {
	scoreCard := ""
	for i := 0; i < 9; i++ {

		scoreCard = playFrame(scoreCard)
	}

	scoreCard = playFinalFrame(scoreCard)
	score := calculateScore(scoreCard)

	return scoreCard, score
}

func playFrame(scoreCard string) string {
	pinsLeft := 10
	firstBowl := rand.Intn(pinsLeft + 1)

	switch firstBowl {
	case 10:
		scoreCard += "X "
		return scoreCard
	case 0:
		scoreCard += "-"
	default:
		scoreCard += strconv.Itoa(firstBowl)
		pinsLeft -= firstBowl
	}

	secondBowl := rand.Intn(pinsLeft + 1)
	if pinsLeft-secondBowl == 0 {
		scoreCard += "/ "
		return scoreCard
	}

	switch secondBowl {
	case 0:
		scoreCard += "- "
	default:
		scoreCard += strconv.Itoa(secondBowl) + " "
	}

	return scoreCard
}

func playFinalFrame(scoreCard string) string {
	pinsLeft := 10
	score := rand.Intn(pinsLeft + 1)
	hasThirdBowl := false

	switch score {
	case 10:
		scoreCard += "X"
		hasThirdBowl = true
		pinsLeft = 10
	case 0:
		scoreCard += "-"
	default:
		scoreCard += strconv.Itoa(score)
		pinsLeft -= score
	}

	score = rand.Intn(pinsLeft + 1)
	// We don't want a spare when the person got a strike before.
	if pinsLeft < 10 && pinsLeft-score == 0 {
		scoreCard += "/ "
		hasThirdBowl = true
		pinsLeft = 10
	} else {
		switch score {
		case 10:
			scoreCard += "X"
			hasThirdBowl = true
			pinsLeft = 10
		case 0:
			scoreCard += "-"
		default:
			scoreCard += strconv.Itoa(score)
			pinsLeft -= score
		}
	}

	if hasThirdBowl {
		score = rand.Intn(pinsLeft + 1)
		if pinsLeft < 10 && pinsLeft-score == 0 {
			scoreCard += "/"
			return scoreCard
		}

		switch score {
		case 10:
			scoreCard += "X"
		case 0:

			scoreCard += "-"
		default:
			scoreCard += strconv.Itoa(score)
		}
	}

	return scoreCard
}

func calculateScore(scoreCard string) int {
	lookahead := func(index, amountToJump int) int {
		if amountToJump < 0 {
			return 0
		}

		score := 0
		for _, char := range scoreCard[index+1:] {
			switch char {
			case 'X', '/':
				score += 10
				amountToJump--
			case '-':
				amountToJump--
			case ' ':
				continue
			default:
				value, _ := strconv.Atoi(string(char))
				score += value
				amountToJump--
			}

			if amountToJump == 0 {
				return score
			}
		}
		return score
	}

	score := 0
	frame := 1
	for i, char := range scoreCard {
		if frame < 10 {
			switch char {
			case 'X':
				score += 10 + lookahead(i, 2)
			case '/':
				score += 10 + lookahead(i, 1)
			case '-':
				continue
			case ' ':
				frame++
			default:
				value, _ := strconv.Atoi(string(char))
				score += value
			}
		} else {
			lengthleft := len(scoreCard) - i
			switch char {
			case 'X':
				score += 10 + lookahead(i, lengthleft)
			case '/':
				score += 10 + lookahead(i, lengthleft-1)
			case '-', ' ':
				continue
			default:
				value, _ := strconv.Atoi(string(char))
				score += value
			}
		}
	}
	return score
}

func readLine() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}
