package application

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoSomething(t *testing.T) {
	// "want" is your expected result
	request := httptest.NewRequest(
		"GET", "/", nil,
	)
	writer := httptest.NewRecorder()

	HomeHandler(writer, request)
	assert.Equal(t, writer.Code, http.StatusOK)

}
