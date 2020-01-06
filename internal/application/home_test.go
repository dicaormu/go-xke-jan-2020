package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDoSomething(t *testing.T) {
	// "want" is your expected result
	request := httptest.NewRequest(
		"GET", "/", nil,
	)
	writer := httptest.NewRecorder()

	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.Info("stating application...")
	HomeHandler(logger)(writer, request)
	assert.Equal(t, writer.Code, http.StatusOK)

}
