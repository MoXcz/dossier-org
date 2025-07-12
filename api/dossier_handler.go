package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MoXcz/dossier-org/db"
	"github.com/MoXcz/dossier-org/internal/database"
	"github.com/MoXcz/dossier-org/models"
)

type DossierHandler struct {
	dossierStore db.DossierStore
}

func NewDossierHandler(dossierStore db.DossierStore) *DossierHandler {
	return &DossierHandler{
		dossierStore: dossierStore,
	}
}

func (h *DossierHandler) HandleGetPostDossier(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var params models.CreateDossierParams
	if err := decoder.Decode(&params); err != nil {
		return APIError{Status: http.StatusBadRequest, Err: err, Msg: "invalid JSON"}
	}

	if errors := params.Validate(); len(errors) > 0 {
		return APIValidateDossierError{Status: http.StatusBadRequest, Errors: errors, Msg: "invalid parameters"}
	}

	dossier, err := models.NewDossierFromParams(params)
	if err != nil {
		return err
	}

	insteredDossier, err := h.dossierStore.CreateDossier(r.Context(), &database.CreateDossierParams{
		Title:      dossier.Title,
		Data:       dossier.Data,
		AssignedTo: dossier.AssignedTo,
	})
	if err != nil {
		return err
	}
	return respJSON(w, http.StatusOK, insteredDossier)
}

func (h *DossierHandler) HandleGetDossiers(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return APIError{Status: http.StatusBadRequest, Msg: "invalid ID", Err: err}
	}

	dossier, err := h.dossierStore.GetDossiersFromUserID(r.Context(), int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return APIError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("invalid ID: there are no users with ID %d", id), Err: err}
	} else if err != nil {
		return err
	}

	return respJSON(w, http.StatusOK, dossier)
}
