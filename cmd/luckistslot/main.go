package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TheGitHubist/LuckistSlot/internal/game"
	"github.com/TheGitHubist/LuckistSlot/internal/storage"
)

// ANSI colors
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Black  = "\033[30m"
)

func main() {
	// Path JSON configurable
	dataPath := os.Getenv("PLAYER_DATA_PATH")
	if dataPath == "" {
		dataPath = "data/player.json"
	}

	store := storage.NewStorage(dataPath)

	// Charger joueur
	p, err := store.LoadPlayerData("admin")
	if err != nil {
		fmt.Println("Erreur chargement joueur :", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== LuckyOps ===")
		fmt.Println("1) Jouer")
		fmt.Println("2) Vérifier solde")
		fmt.Println("3) Quitter")
		fmt.Print("Choix : ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Mise en jetons : ")
			scanner.Scan()
			stakeInput := scanner.Text()
			stake, err := strconv.Atoi(strings.TrimSpace(stakeInput))
			if err != nil || stake <= 0 {
				fmt.Println("Mise invalide.")
				continue
			}
			if p.Tokens < stake {
				fmt.Println("Vous n'avez pas assez de jetons.")
				continue
			}

			// Jouer avec animation
			result := spinWithAnimation(stake)

			// Afficher rouleaux finaux
			printReels(result.Reels)

			// Afficher résultat
			fmt.Println(result.Message)

			// Mettre à jour jetons
			p.Tokens -= stake
			p.Tokens += result.WinAmount

			// Sauvegarder
			if err := store.SavePlayerData(p); err != nil {
				fmt.Println("Erreur sauvegarde :", err)
			}

		case "2":
			fmt.Printf("Solde actuel : %d jetons\n", p.Tokens)

		case "3":
			fmt.Println("Merci d'avoir joué ! À bientôt !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// spinWithAnimation fait tourner chaque rouleau individuellement
func spinWithAnimation(stake int) game.Result {
	reels := make([]game.Symbol, 5)
	frames := 10
	for f := 0; f < frames; f++ {
		for i := 0; i < 5; i++ {
			reels[i] = game.Symbols[rand.Intn(len(game.Symbols))]
		}
		printReels(reels)
		time.Sleep(100 * time.Millisecond)
		clearPreviousLines(3)
	}

	// Tirage final réel
	return game.Spin(stake)
}

// printReels affiche les rouleaux ASCII alignés
func printReels(reels []game.Symbol) {
	// Ligne supérieure
	fmt.Print("+")
	for range reels {
		fmt.Print("-------+")
	}
	fmt.Println()

	// Ligne des symboles centrés
	fmt.Print("|")
	for _, s := range reels {
		color := getColor(s)
		sym := s.Name
		// centrer dans 7 caractères
		fmt.Printf(" %s%-5s%s |", color, sym, Reset)
	}
	fmt.Println()

	// Ligne inférieure
	fmt.Print("+")
	for range reels {
		fmt.Print("-------+")
	}
	fmt.Println()
}

// getColor renvoie couleur ANSI
func getColor(s game.Symbol) string {
	switch s.Color {
	case "Hearts", "Diamonds":
		return Red
	case "None":
		return Yellow
	default:
		return Black
	}
}

// clearPreviousLines efface n lignes pour animation
func clearPreviousLines(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\033[F\033[K")
	}
}
