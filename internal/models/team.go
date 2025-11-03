package models

type Country struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:120"`
}

type Team struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:160"`
	ShortCode string `gorm:"size:16"`
	CountryID uint
	RegionID  uint
}
