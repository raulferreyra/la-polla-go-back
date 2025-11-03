package server

import (
	"go-backend/internal/router"

	"github.com/gin-gonic/gin"
)

func New(deps router.Deps) *gin.Engine {
	return router.Setup(deps)
}
