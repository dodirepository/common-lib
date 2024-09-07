package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

// Render :nodoc:
func Render(response interface{}, HTTPStatusCode int, writer http.ResponseWriter) {
	if response == nil {
		writer.WriteHeader(HTTPStatusCode)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(HTTPStatusCode)
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Print(err.Error())
	}
}
