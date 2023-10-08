package server

import (
	"net/http"

	"github.com/go-cinch/layout/internal/service"
)

func HealthHandler(svc *service.GameService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/pub/healthcheck", svc.HealthCheck)
	return mux
}
