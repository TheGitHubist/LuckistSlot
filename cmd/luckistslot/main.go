package main

import (
	"fmt"

	"github.com/TheGitHubist/LuckistSlot/internal/player"
)

func main() {
	// Example usage of the Player struct and its methods
	player1 := player.NewPlayer("Alice", 100)
	fmt.Printf("Player: %s, Tokens: %d\n", player1.Name, player1.Tokens)

	// Attempt to claim daily reward
	if player1.ClaimDailyReward(50, 24) {
		fmt.Printf("Daily reward claimed! New token count: %d\n", player1.Tokens)
	} else {
		fmt.Println("Cannot claim daily reward yet.")
	}
}
