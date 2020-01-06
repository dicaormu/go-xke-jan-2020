package application

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HomeHandler(logger *log.Logger) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Info("handling")
		writer.WriteHeader(http.StatusOK)
	}
}
