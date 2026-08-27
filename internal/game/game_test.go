package game

import "testing"

func TestSpinLength(t *testing.T) {
	result := Spin(10)
	if len(result.Reels) != 5 {
		t.Errorf("Spin doit retourner 5 symboles, got %d", len(result.Reels))
	}
}

func TestEvaluateHandMaxJackpot(t *testing.T) {
	reels := []Symbol{
		{"Joker", "Joker", "None", "joker"},
		{"Joker", "Joker", "None", "joker"},
		{"Joker", "Joker", "None", "joker"},
		{"Joker", "Joker", "None", "joker"},
		{"Joker", "Joker", "None", "joker"},
	}
	winType := evaluateHand(reels)
	if winType != "Max Jackpot" {
		t.Errorf("Expected Max Jackpot, got %s", winType)
	}
}

func TestEvaluateHandGreatJackpot(t *testing.T) {
	reels := []Symbol{
		{"A♥", "A", "Hearts", "card"},
		{"A♦", "A", "Diamonds", "card"},
		{"A♣", "A", "Clubs", "card"},
		{"A♠", "A", "Spades", "card"},
		{"A♥", "A", "Hearts", "card"},
	}
	winType := evaluateHand(reels)
	if winType != "Great Jackpot" {
		t.Errorf("Expected Great Jackpot, got %s", winType)
	}
}

func TestEvaluateHandQuinteFlush(t *testing.T) {
	reels := []Symbol{
		{"10♥", "10", "Hearts", "card"},
		{"J♥", "J", "Hearts", "card"},
		{"Q♥", "Q", "Hearts", "card"},
		{"K♥", "K", "Hearts", "card"},
		{"A♥", "A", "Hearts", "card"},
	}
	winType := evaluateHand(reels)
	if winType != "Royal Flush" {
		t.Errorf("Expected Royal Flush, got %s", winType)
	}
}

func TestCalculateWinBrelan(t *testing.T) {
	reels := []Symbol{
		{"10♥", "10", "Hearts", "card"},
		{"10♦", "10", "Diamonds", "card"},
		{"10♣", "10", "Clubs", "card"},
		{"J♠", "J", "Spades", "card"},
		{"Q♥", "Q", "Hearts", "card"},
	}
	win := calculateWinAmount(reels, 10, "Win")
	if win <= 0 {
		t.Errorf("Brelan doit générer un gain > 0, got %d", win)
	}
}
