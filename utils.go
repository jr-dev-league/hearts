package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func encodeResponse(w http.ResponseWriter, req *http.Request, i interface{}, code int) ([]byte, error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(bytes)
	}

	return bytes, err
}
