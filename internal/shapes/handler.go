package shapes

import (
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/shapes/ports"
	"github.com/ExonegeS/AP1_Assignment1/internal/shapes/service"
	"github.com/gorilla/mux"
)

func NewHandler(s service.Service, logger *slog.Logger) http.Handler {
	r := mux.NewRouter()
	r.Handle("/shapes/print", ports.NewEndpointPrintShapes(s, logger)).Methods(http.MethodGet)

	// Register other routes as needed
	return r
}
