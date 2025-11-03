package server

import (
	"go-backend/internal/router"
)

func New(deps router.Deps) *router.Engine {
	return router.Setup(deps)
}
