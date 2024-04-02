package routes

import (
	"net/http"

	"github.com/0x30c4/Go-Backend-BoilerPlate/internal/handlers"
)

func Router() http.Handler {
  mux := &http.ServeMux{}

  mux.HandleFunc("GET /", handlers.Index)

  return mux
}
