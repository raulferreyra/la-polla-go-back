package server

import (
	"polla/internal/router"
)

func New(deps router.Deps) *router.Engine {
	return router.Setup(deps)
}
