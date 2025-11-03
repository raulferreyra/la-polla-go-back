package models

import "time"

type Tournament struct {
	ID        uint   `gorm:"primaryKey"`
	CompanyID uint   `gorm:"index"`    // habilitado por empresa
	Name      string `gorm:"size:160"` // "Mundial 2026", "Libertadores 2025"
	Type      string `gorm:"size:40"`  // Mundial/Libertadores/Clausura...
	Enabled   bool
	Phase     string `gorm:"size:20"` // BETTING|SCORING
	CreatedAt time.Time
	UpdatedAt time.Time
}
