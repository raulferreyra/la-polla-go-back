package models

import "time"

type Match struct {
	ID           uint `gorm:"primaryKey"`
	TournamentID uint `gorm:"index"`
	HomeTeamID   uint `gorm:"index"`
	AwayTeamID   uint `gorm:"index"`
	Kickoff      time.Time
	Stage        string `gorm:"size:60"` // grupos, 8vos, final
	// Resultado real:
	ResultType string `gorm:"size:30"` // HOME|AWAY|DRAW|HOME_PEN|AWAY_PEN|HOME_ET|AWAY_ET
	HomeGoals  *int
	AwayGoals  *int
}
