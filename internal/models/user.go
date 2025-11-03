package models

import "time"

type Role string

const (
	RoleSuperAdmin Role = "SUPERADMIN"
	RoleGroupAdmin Role = "GROUP_ADMIN"
	RoleUser       Role = "USER"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	CompanyID uint   `gorm:"index"`
	EmailEnc  string `gorm:"type:text"`      // cifrado AES (email sensible)
	EmailHash string `gorm:"index;size:190"` // para búsquedas (sha256/email)
	Password  string `gorm:"size:120"`       // bcrypt hash
	Role      Role   `gorm:"size:20"`
	Display   string `gorm:"size:120"`
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
