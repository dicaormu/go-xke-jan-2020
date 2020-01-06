package diagnostics

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ReadinessHandler(logger *log.Logger) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("is ready")
		writer.WriteHeader(http.StatusOK)
	}
}

func LivenessHandler(logger *log.Logger) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("is alive!")
		writer.WriteHeader(http.StatusOK)
	}
}
