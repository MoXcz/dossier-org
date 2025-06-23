package api

import (
	"net/http"

	"github.com/MoXcz/dossier-org/helpers"
)

type appHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Make(handler appHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			switch e := err.(type) {
			case APIError:
				respJSON(w, e.Status, map[string]any{"error": e.Msg})
				helpers.Logger.Error("Server error", "err", e.Err, "status", e.Status)
			case APIValidateUserError:
				respJSON(w, e.Status, map[string]any{"error": e.Errors})
				helpers.Logger.Error("Validate user error", "err", e.Msg, "status", e.Status)
			default:
				respJSON(w, http.StatusInternalServerError, "Internal Server Error")
				helpers.Logger.Error("Validate user error", "err", err)
			}
		}
	}
}

type APIError struct {
	Status int
	Msg    string
	Err    error
}

func (err APIError) Error() string {
	return err.Msg
}

// TODO: Adjust this, Errors is more like Msg because it's on client-side
type APIValidateUserError struct {
	Status int
	Msg    string
	Errors map[string]string
}

func (err APIValidateUserError) Error() string {
	return err.Msg
}
