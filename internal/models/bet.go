package models

import "time"

type Bet struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index"`
	MatchID     uint   `gorm:"index"`
	Pick        string `gorm:"size:30"` // HOME|AWAY|DRAW|HOME_PEN|...
	Points      int
	SubmittedAt time.Time
}
