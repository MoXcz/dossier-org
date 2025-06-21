package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

var Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func respJSON(w http.ResponseWriter, code int, jsonResp any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.Marshal(jsonResp)
	if err != nil {
		return err
	}
	w.Write(data)

	return nil
}
