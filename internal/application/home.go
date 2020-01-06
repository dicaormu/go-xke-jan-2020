package application

import (
	"fmt"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("handling")
	writer.WriteHeader(http.StatusOK)
}
