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

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetPostUser(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var params models.CreateUserParams
	err := decoder.Decode(&params)
	if err != nil {
		return APIError{Status: http.StatusBadRequest, Err: err, Msg: "invalid JSON"}
	}

	if errors := params.Validate(); len(errors) > 0 {
		return APIValidateUserError{Status: http.StatusBadRequest, Errors: errors, Msg: "invalid parameters"}
	}

	user, err := models.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insteredUser, err := h.userStore.CreateUser(r.Context(), &database.CreateUserParams{
		Name:              user.Name,
		Email:             user.Email,
		Encryptedpassword: user.EncryptedPassword,
	})
	if err != nil {
		return err
	}
	return respJSON(w, http.StatusOK, insteredUser)
}

func (h *UserHandler) HandleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.userStore.GetUsers(r.Context())
	if err != nil {
		return err
	}
	return respJSON(w, http.StatusOK, users)
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    "invalid ID",
			Err:    err,
		}
	}

	user, err := h.userStore.GetUserByID(r.Context(), int32(id))
	if errors.Is(err, sql.ErrNoRows) {
		return APIError{
			Status: http.StatusBadRequest,
			Msg:    fmt.Sprintf("invalid ID: there are no users with ID %d", id),
			Err:    err,
		}
	} else if err != nil {
		return err
	}

	return respJSON(w, http.StatusOK, user)
}
