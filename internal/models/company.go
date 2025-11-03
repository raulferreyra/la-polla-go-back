package models

import "time"

type Company struct {
	ID        uint   `gorm:"primaryKey"`
	Slug      string `gorm:"uniqueIndex;size:64"`
	Name      string `gorm:"size:120"`
	LogoURL   string `gorm:"size:255"`
	ThemeJSON string `gorm:"type:text"`
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
