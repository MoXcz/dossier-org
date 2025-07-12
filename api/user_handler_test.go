package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MoXcz/dossier-org/db"
	"github.com/MoXcz/dossier-org/helpers"
	"github.com/MoXcz/dossier-org/models"
	_ "github.com/lib/pq"
)

type testdb struct {
	db.UserStore
}

const (
	testdburi = "postgres://postgres:password@localhost:5432/dossier?sslmode=disable"
)

func setup() *testdb {
	sqlDB, err := db.OpenDB(testdburi)
	if err != nil {
		helpers.Logger.Error(err.Error())
		os.Exit(1)
	}

	return &testdb{
		UserStore: db.NewPostgresUserStore(sqlDB),
	}
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup()
	defer tdb.teardown(t)

	mux := http.NewServeMux()
	userHandler := NewUserHandler(tdb.UserStore)

	mux.HandleFunc("POST /", Make(userHandler.HandleGetPostUser))

	params := models.CreateUserParams{
		Name:     "Jane",
		Email:    "jane@mail.com",
		Password: "Jane12345!",
		RoleID:   1,
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	var user models.User
	json.NewDecoder(rr.Body).Decode(&user)
	if user.ID < 0 {
		t.Error("expecting a valid user id to be set")
	}
	if len(user.HashPassword) > 0 {
		t.Error("expected encrypted password to not be included in response")
	}
	if user.Name != params.Name {
		t.Errorf("expected first name %s but got %s", params.Name, user.Name)
	}
	if user.Email != params.Email {
		t.Errorf("expected first name %s but got %s", params.Email, user.Email)
	}
	if user.RoleID != params.RoleID {
		t.Errorf("expected role id %d but got %d", params.RoleID, user.RoleID)
	}

	fmt.Println(user)
	fmt.Println(rr.Code)
}
