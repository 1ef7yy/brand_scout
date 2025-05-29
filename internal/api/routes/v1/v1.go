package v1

import (
	"github.com/1ef7yy/brand_scout/internal/view"
	"github.com/1ef7yy/brand_scout/pkg/logger"
)

type Router struct {
	View view.ViewIFace
	log  logger.Logger
}

func NewRouter(view view.ViewIFace, log logger.Logger) *Router {
	return &Router{
		View: view,
		log:  log,
	}
}
