package player

import (
	"time"
)

type Player struct {
	Name            string    `json:"name"`
	Tokens          int       `json:"tokens"`
	LastDailyReward time.Time `json:"lastDailyReward"`
}

func NewPlayer(name string, InitialTokens int) *Player {
	return &Player{
		Name:            name,
		Tokens:          InitialTokens,
		LastDailyReward: time.Time{},
	}
}

func (p *Player) CanClaimDailyReward(cooldownInHours int) bool {
	if p.LastDailyReward.IsZero() {
		return true
	}

	nextClaimTime := p.LastDailyReward.Add(time.Duration(cooldownInHours) * time.Hour)
	return time.Now().After(nextClaimTime)
}

func (p *Player) ClaimDailyReward(rewardAmount int, cooldownInHours int) bool {
	if !p.CanClaimDailyReward(cooldownInHours) {
		return false
	}
	p.Tokens += rewardAmount
	p.LastDailyReward = time.Now()
	return true
}

func (p *Player) AddTokens(amount int) {
	p.Tokens += amount
}

func (p *Player) RemoveTokens(amount int) bool {
	if p.Tokens < amount {
		return false
	}
	p.Tokens -= amount
	return true
}
