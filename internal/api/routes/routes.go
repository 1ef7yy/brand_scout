package routes

import (
	"net/http"

	v1 "github.com/1ef7yy/brand_scout/internal/api/routes/v1"
	"github.com/1ef7yy/brand_scout/internal/view"
	"github.com/1ef7yy/brand_scout/pkg/logger"
)

func InitHandlers(view view.ViewIFace) *http.ServeMux {
	mux := http.NewServeMux()
	v1 := v1.NewRouter(view, logger.NewLogger())

	mux.Handle("/", v1.Quotes())

	return mux
}
