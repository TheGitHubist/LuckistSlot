package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/TheGitHubist/LuckistSlot/internal/player"
)

type Storage struct {
	FilePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{FilePath: filePath}
}

func (s *Storage) SavePlayerData(p *player.Player) error {
	file, err := os.Create(s.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(p)
}

func (s *Storage) LoadPlayerData(name string) (*player.Player, error) {
	if _, err := os.Stat(s.FilePath); errors.Is(err, os.ErrNotExist) {
		return player.NewPlayer(name, 100), nil
	}

	file, err := os.Open(s.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var p player.Player
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}
