package middleware

import (
	"net/http"
	"strings"

	"polla/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Tenant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// asume ruta: /:companySlug/...
		parts := strings.Split(strings.Trim(c.Request.URL.Path, "/"), "/")
		if len(parts) == 0 {
			c.Next()
			return
		}
		slug := parts[0]
		var comp models.Company
		if err := db.Where("slug = ? AND enabled = true", slug).First(&comp).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "empresa no encontrada"})
			return
		}
		c.Set("tenant_id", comp.ID)
		c.Set("tenant_slug", comp.Slug)
		c.Next()
	}
}
