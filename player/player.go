package player

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Player struct {
	ScoreCard   string
	Score       int
	Personality string
	Skill       int
	Mentality   int
}

func (p *Player) PlayBowl(pinsLeft int) int {
	// lowerHalf := math.Ceil(float64(pinsLeft / 2))
	averageStatus := p.Skill + p.Mentality
	weighting := 10 - averageStatus

	fmt.Printf("DEBUG: weighting %v\n", weighting)
	pins := make(map[int]int, pinsLeft)
	for x := 0; x <= pinsLeft; x++ {
		pins[x] = 10
	}

	totalWeight := 0
	for _, weight := range pins {
		totalWeight += weight
	}

	r := rand.Intn(totalWeight)
	for i := 0; i < pinsLeft; i++ {
		r -= pins[i]
		if r <= 0 {
			return i
		}
	}

	return 0
}

func (p *Player) PlayFrame() {
	pinsLeft := 10
	firstBowl := p.PlayBowl(pinsLeft + 1)

	switch firstBowl {
	case 10:
		p.ScoreCard += "X "
		return
	case 0:
		p.ScoreCard += "-"
	default:
		p.ScoreCard += strconv.Itoa(firstBowl)
		pinsLeft -= firstBowl
	}

	secondBowl := p.PlayBowl(pinsLeft + 1)
	if pinsLeft-secondBowl == 0 {
		p.ScoreCard += "/ "
		return
	}

	switch secondBowl {
	case 0:
		p.ScoreCard += "- "
	default:
		p.ScoreCard += strconv.Itoa(secondBowl) + " "
	}
}

func (p *Player) PlayFinalFrame() {
	pinsLeft := 10
	score := p.PlayBowl(pinsLeft + 1)
	hasThirdBowl := false

	switch score {
	case 10:
		p.ScoreCard += "X"
		hasThirdBowl = true
		pinsLeft = 10
	case 0:
		p.ScoreCard += "-"
	default:
		p.ScoreCard += strconv.Itoa(score)
		pinsLeft -= score
	}

	score = p.PlayBowl(pinsLeft + 1)
	// We don't want a spare when the person got a strike before.
	if pinsLeft < 10 && pinsLeft-score == 0 {
		p.ScoreCard += "/ "
		hasThirdBowl = true
		pinsLeft = 10
	} else {
		switch score {
		case 10:
			p.ScoreCard += "X"
			hasThirdBowl = true
			pinsLeft = 10
		case 0:
			p.ScoreCard += "-"
		default:
			p.ScoreCard += strconv.Itoa(score)
			pinsLeft -= score
		}
	}

	if hasThirdBowl {
		score = p.PlayBowl(pinsLeft + 1)
		if pinsLeft < 10 && pinsLeft-score == 0 {
			p.ScoreCard += "/"
		}

		switch score {
		case 10:
			p.ScoreCard += "X"
		case 0:
			p.ScoreCard += "-"
		default:
			p.ScoreCard += strconv.Itoa(score)
		}
	}
}

func (p *Player) CalculateScore() {
	lookahead := func(index, amountToJump int) int {
		if amountToJump < 0 {
			return 0
		}

		score := 0
		for _, char := range p.ScoreCard[index+1:] {
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

	frame := 1
	for i, char := range p.ScoreCard {
		if frame < 10 {
			switch char {
			case 'X':
				p.Score += 10 + lookahead(i, 2)
			case '/':
				p.Score += 10 + lookahead(i, 1)
			case '-':
				continue
			case ' ':
				frame++
			default:
				value, _ := strconv.Atoi(string(char))
				p.Score += value
			}
		} else {
			lengthleft := len(p.ScoreCard) - i
			switch char {
			case 'X':
				p.Score += 10 + lookahead(i, lengthleft)
			case '/':
				p.Score += 10 + lookahead(i, lengthleft-1)
			case '-', ' ':
				continue
			default:
				value, _ := strconv.Atoi(string(char))
				p.Score += value
			}
		}
	}
}

func NewPlayer() *Player {
	p := new(Player)
	p.Skill = 5
	p.Mentality = 5
	return p
}
