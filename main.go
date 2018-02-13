package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	player "github.com/omgitsotis/bowling/player"
)

func main() {
	rand.Seed(time.Now().Unix())
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
	humanPlayer := player.NewPlayer()
	cpuPlayer := player.NewPlayer()

	frameCount := 1
	for {
		fmt.Printf("Frame %d\n", frameCount)
		humanPlayer.PlayFrame()
		fmt.Printf("player score so far [%s]\n", humanPlayer.ScoreCard)
		cpuPlayer.PlayFrame()
		fmt.Printf("cpu score card so far[%s]\n", cpuPlayer.ScoreCard)
		cpuPlayer.ChangeStats(player.Skill, -1)
		fmt.Printf("cpu player skill [%d]\n", cpuPlayer.Skill)
		fmt.Print("Continue? ")
		line := readLine()
		switch strings.ToLower(line) {
		case "yes", "y":
			frameCount++
		default:
			os.Exit(0)
		}

		if frameCount > 9 {
			break
		}
	}

	humanPlayer.PlayFinalFrame()
	fmt.Printf("final player score [%s]\n", humanPlayer.ScoreCard)
	cpuPlayer.ChangeStats(player.Skill, 9)
	fmt.Printf("cpu player skill [%d]\n", cpuPlayer.Skill)
	cpuPlayer.PlayFinalFrame()
	fmt.Printf("final cpu score [%s]\n", cpuPlayer.ScoreCard)

	humanPlayer.CalculateScore()
	cpuPlayer.CalculateScore()

	var finalScore string
	if humanPlayer.Score > cpuPlayer.Score {
		finalScore = fmt.Sprintf("player wins [%d] to [%d]\n", humanPlayer.Score, cpuPlayer.Score)
	} else {
		finalScore = fmt.Sprintf("cpu wins [%d] to [%d]\n", cpuPlayer.Score, humanPlayer.Score)
	}

	fmt.Printf(finalScore)
}

func readLine() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}
