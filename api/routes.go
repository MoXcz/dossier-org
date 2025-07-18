package api

import (
	"net/http"
	"os"

	"github.com/MoXcz/dossier-org/db"
	"github.com/MoXcz/dossier-org/helpers"
	"github.com/joho/godotenv"
)

func Routes() http.Handler {
	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		helpers.Logger.Info("DB_URL variable must be defined (use a .env file)", "DB_URL", dbURL)
	}

	sqlDB, err := db.OpenDB(dbURL)
	if err != nil {
		helpers.Logger.Error(err.Error())
		os.Exit(1)
	}

	var (
		userHandler    = NewUserHandler(db.NewPostgresUserStore(sqlDB))
		dossierHandler = NewDossierHandler(db.NewPostgresDossierStore(sqlDB))
		mux            = http.NewServeMux()
	)

	// users
	mux.HandleFunc("GET /user/{id}", Make(userHandler.HandleGetUser))
	mux.HandleFunc("GET /user", Make(userHandler.HandleGetUsers))
	mux.HandleFunc("POST /user", Make(userHandler.HandleGetPostUser))

	// dossiers
	mux.HandleFunc("GET /user/{id}/dossier", Make(dossierHandler.HandleGetDossiers))
	mux.HandleFunc("POST /dossier", Make(dossierHandler.HandleGetPostDossier))

	return mux
}
