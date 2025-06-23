package api

import (
	"encoding/json"
	"net/http"
)

func respJSON(w http.ResponseWriter, code int, jsonPayload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.Marshal(jsonPayload)
	if err != nil {
		return err
	}
	w.Write(data)

	return nil
}
