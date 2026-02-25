package game

import (
	"fmt"
	"math/rand"
	"time"
)

var Symbols = []Symbol{
	{"10 ♠", "10", "black", "card"},
	{"J ♠", "J", "black", "card"},
	{"Q ♠", "Q", "black", "card"},
	{"K ♠", "K", "black", "card"},
	{"A ♠", "A", "black", "card"},

	{"10 ♥", "10", "red", "card"},
	{"J ♥", "J", "red", "card"},
	{"Q ♥", "Q", "red", "card"},
	{"K ♥", "K", "red", "card"},
	{"A ♥", "A", "red", "card"},

	{"10 ♦", "10", "red", "card"},
	{"J ♦", "J", "red", "card"},
	{"Q ♦", "Q", "red", "card"},
	{"K ♦", "K", "red", "card"},
	{"A ♦", "A", "red", "card"},

	{"10 ♣", "10", "black", "card"},
	{"J ♣", "J", "black", "card"},
	{"Q ♣", "Q", "black", "card"},
	{"K ♣", "K", "black", "card"},
	{"A ♣", "A", "black", "card"},

	{"Joker", "Joker", "None", "joker"},
}

func Spin(stake int) Result {
	rand.Seed(time.Now().UnixNano())
	reels := make([]Symbol, 5)
	for i := 0; i < 5; i++ {
		reels[i] = Symbols[rand.Intn(len(Symbols))]
	}

	winType := evaluateHand(reels)
	winAmount := calculateWinAmount(reels, stake, winType)

	return Result{
		Reels:     reels,
		WinAmount: winAmount,
		WinType:   winType,
		Message:   generateMessage(winType, winAmount),
	}
}

func evaluateHand(reels []Symbol) string {
	jokerCount := 0
	symbolCount := map[string]int{}
	colorCount := map[string]int{}

	for _, symbol := range reels {
		if symbol.Type == "joker" {
			jokerCount++
			continue
		}
		symbolCount[symbol.Value]++
		colorCount[symbol.Color]++
	}

	if jokerCount == 5 {
		return "Max Jackpot"
	}

	for _, count := range symbolCount {
		if count+jokerCount == 5 {
			return "Great Jackpot"
		}
	}

	// testing royal flush

	valuesNeeded := map[string]bool{"10": false, "J": false, "Q": false, "K": false, "A": false}
	for _, symbol := range reels {
		if symbol.Type == "card" {
			valuesNeeded[symbol.Value] = true
		}
	}

	for color, count := range colorCount {
		if count == 5 {
			allFound := true
			for val := range valuesNeeded {
				found := false
				for _, symbol := range reels {
					if symbol.Value == val && symbol.Color == color {
						found = true
						break
					}
				}
				if !found {
					allFound = false
					break
				}
			}
			if allFound {
				return "Royal Flush"
			}
		}
	}

	for _, count := range symbolCount {
		if count >= 3 {
			return "Win"
		}
	}

	return "No Win"
}

func calculateWinAmount(reels []Symbol, stake int, winType string) int {
	baseWin := stake * 3
	switch winType {
	case "Max Jackpot":
		return stake * 999
	case "Royal Flush":
		return stake * 200
	case "Great Jackpot":
		return stake * 32
	case "Win":
		symbolCount := map[string]int{}
		colorCount := map[string]int{}
		for _, symbol := range reels {
			if symbol.Type != "joker" {
				symbolCount[symbol.Value]++
				colorCount[symbol.Color]++
			}
		}
		winMultiplier := 1.25
		for value, count := range symbolCount {
			if count >= 3 {
				switch value {
				case "10":
					winMultiplier = 1.25
				case "J":
					winMultiplier = 1.33
				case "Q":
					winMultiplier = 1.5
				case "K":
					winMultiplier = 1.66
				case "A":
					winMultiplier = 1.75
				}
			}
		}
		return int(float64(baseWin) * winMultiplier)
	default:
		return 0
	}
}

func generateMessage(winType string, winAmount int) string {
	switch winType {
	case "Max Jackpot":
		return fmt.Sprintf("Congratulations! You hit the Max Jackpot and won %d coins!", winAmount)
	case "Royal Flush":
		return fmt.Sprintf("Amazing! You got a Royal Flush and won %d coins!", winAmount)
	case "Great Jackpot":
		return fmt.Sprintf("Great job! You hit the Great Jackpot and won %d coins!", winAmount)
	case "Win":
		return fmt.Sprintf("You won %d coins!", winAmount)
	default:
		return "Better luck next time!"
	}
}
