package main

import (
	// "bufio"
	// "os"

	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	// score, _ := readScore("X X X X X X X X X XXX")
	// fmt.Println(score)
	startGame()
}

func startGame() {
	rand.Seed(time.Now().Unix())
	scoreCard := ""
	for i := 0; i < 9; i++ {
		fmt.Printf("Frame %d\n", i+1)
		scoreCard = playFrame(scoreCard)
	}

	fmt.Println("Frame 10")
	scoreCard = playFinalFrame(scoreCard)
	fmt.Println(scoreCard)
	calculateScore(scoreCard)
}

func playFrame(scoreCard string) string {
	pinsLeft := 10
	firstBowl := rand.Intn(pinsLeft + 1)

	switch firstBowl {
	case 10:
		fmt.Println("Bowl 1 [X]")
		scoreCard += "X "
		return scoreCard
	case 0:
		fmt.Println("Bowl 1 [-]")
		scoreCard += "-"
	default:
		fmt.Printf("Bowl 1 [%d]\n", firstBowl)
		scoreCard += strconv.Itoa(firstBowl)
		pinsLeft -= firstBowl
	}

	secondBowl := rand.Intn(pinsLeft + 1)
	if pinsLeft-secondBowl == 0 {
		fmt.Println("Bowl 2 [/]")
		scoreCard += "/ "
		return scoreCard
	}

	switch secondBowl {
	case 0:
		fmt.Println("Bowl 2 [-]")
		scoreCard += "- "
	default:
		fmt.Printf("Bowl 2 [%d]\n", secondBowl)
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
		fmt.Println("Bowl 1 [X]")
		scoreCard += "X"
		hasThirdBowl = true
		pinsLeft = 10
	case 0:
		fmt.Println("Bowl 1 [-]")
		scoreCard += "-"
	default:
		fmt.Printf("Bowl 1 [%d]\n", score)
		scoreCard += strconv.Itoa(score)
		pinsLeft -= score
	}

	score = rand.Intn(pinsLeft + 1)
	// We don't want a spare when the person got a strike before.
	if pinsLeft < 10 && pinsLeft-score == 0 {
		fmt.Println("Bowl 2 [/]")
		scoreCard += "/ "
		hasThirdBowl = true
		pinsLeft = 10
	} else {
		switch score {
		case 10:
			fmt.Println("Bowl 2 [X]")
			scoreCard += "X"
			hasThirdBowl = true
			pinsLeft = 10
		case 0:
			fmt.Println("Bowl 2 [-]")
			scoreCard += "-"
		default:
			fmt.Printf("Bowl 2 [%d]\n", score)
			scoreCard += strconv.Itoa(score)
			pinsLeft -= score
		}
	}

	if hasThirdBowl {
		score = rand.Intn(pinsLeft + 1)
		if pinsLeft < 10 && pinsLeft-score == 0 {
			fmt.Println("Bowl 3 [/]")
			scoreCard += "/"
			return scoreCard
		}

		switch score {
		case 10:
			fmt.Println("Bowl 3 [X]")
			scoreCard += "X"
		case 0:
			fmt.Println("Bowl 3 [-]")
			scoreCard += "-"
		default:
			fmt.Printf("Bowl 3 [%d]\n", score)
			scoreCard += strconv.Itoa(score)
		}
	}

	return scoreCard
}

calculateScore(scoreCard string) int {
	score =
	for i, char := range scores {
		switch char {
		case 'X':
			s
		}
	}
}

lookahead(index, amountToJump int) int{
	
}
