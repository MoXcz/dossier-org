package models

import (
	"encoding/json"
)

// server-side parameters
type CreateDossierParams struct {
	Title      string          `json:"title"`
	Data       json.RawMessage `json:"data"`
	AssignedTo int64           `json:"assigned_to"`
}

func (params CreateDossierParams) Validate() map[string]string {
	errors := map[string]string{}
	// if len(params.Title) < minNameLen {
	// }
	// if len(params.Data) < minPasswordLen {
	// 	errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	// }
	// if params.AssignedTo != 1 {
	// 	errors["role_id"] = "role is not valid"
	// }
	return errors
}

type Dossier struct {
	Title      string          `json:"title"`
	Data       json.RawMessage `json:"data"`
	AssignedTo int64           `json:"assigned_to"`
}

func NewDossierFromParams(params CreateDossierParams) (*Dossier, error) {
	return &Dossier{
		Title:      params.Title,
		Data:       params.Data,
		AssignedTo: params.AssignedTo,
	}, nil
}
