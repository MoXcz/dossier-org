package api

import (
	"net/http"

	"github.com/MoXcz/dossier-org/helpers"
)

type appHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Make(handler appHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			helpers.Logger.Error("")
			if err, ok := err.(APIServerError); ok {
				respJSON(w, err.Status, err)
				helpers.Logger.Error("Server error", "err", err, "status", err.Status)
			}
			if err, ok := err.(APIValidateUserError); ok {
				respJSON(w, err.Status, err)
				helpers.Logger.Error("Validate user error", "err", err, "status", err.Status)
			}
		}

		// should not happen, left it here "just in case" while debugging
		if err := handler(w, r); err != nil {
			helpers.Logger.Error("handler error", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
