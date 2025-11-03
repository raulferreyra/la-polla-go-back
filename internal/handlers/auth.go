package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"go-backend/internal/auth"
	"go-backend/internal/models"
)

type LoginReq struct {
	Company string `json:"company"` // redundante si viene por path, útil para /auth/login general
	Email   string `json:"email"`
	Pass    string `json:"password"`
}

func Login(db *gorm.DB, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in LoginReq
		if err := c.BindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input"})
			return
		}
		// resolver empresa del contexto si viene por path
		tenantIDAny, _ := c.Get("tenant_id")
		var tenantID uint
		if tenantIDAny != nil {
			tenantID = tenantIDAny.(uint)
		}

		eh := sha256.Sum256([]byte(in.Email))
		hash := hex.EncodeToString(eh[:])

		var u models.User
		q := db.Where("email_hash = ? AND active = true", hash)
		if tenantID > 0 {
			q = q.Where("company_id = ?", tenantID)
		}
		if err := q.First(&u).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales"})
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Pass)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales"})
			return
		}
		tok, _ := auth.Sign(jwtSecret, u.ID, u.CompanyID, string(u.Role))
		c.JSON(200, gin.H{"token": tok, "role": u.Role})
	}
}
