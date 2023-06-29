package api

import (
	"encoding/json"
	"net/http"
	"statuarius/internal/values"
)

type Handler func(w http.ResponseWriter, r *http.Request) *ServerResponse

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := h(w, r)
	responseByte, err := json.Marshal(response)
	if err != nil {
		writeErrorResponse(w, err, values.Error, "unable to marshal server response")
		return
	}
	writeJSONResponse(w, responseByte, response.StatusCode)
}
