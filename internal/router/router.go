package router

import (
	"go-backend/internal/handlers"
	"go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Deps struct {
	DB        *gorm.DB
	JWTSecret string
}

func Setup(d Deps) *gin.Engine {
	r := gin.Default()
	r.GET("/health", handlers.Health)

	// rutas por empresa
	company := r.Group("/:company")
	company.Use(middleware.Tenant(d.DB))
	{
		company.POST("/auth/login", handlers.Login(d.DB, d.JWTSecret))
		company.POST("/admin/roster/upload", middleware.RequireAuth(d.JWTSecret, "GROUP_ADMIN"), handlers.AdminRosterUpload)
		// TODO: /tournaments, /matches, /bets ...
	}

	return r
}
