package models

type Region struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:120"` // CONMEBOL, UEFA, CONCACAF
}
