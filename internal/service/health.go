package service

import (
	"net/http"

	"github.com/go-cinch/common/log"
)

func (*GameService) HealthCheck(writer http.ResponseWriter, _ *http.Request) {
	log.Debug("healthcheck")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("{}"))
	return
}
