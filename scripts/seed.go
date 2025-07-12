package main

import (
	"context"
	"os"

	"github.com/MoXcz/dossier-org/db"
	"github.com/MoXcz/dossier-org/helpers"
	"github.com/MoXcz/dossier-org/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	ctx = context.Background()
)

func seedRoles(dbQueries *database.Queries) {
	roles := []database.Role{
		{
			Name:        "admin",
			Description: "administrator user",
		},
		{
			Name:        "employee",
			Description: "normal employee",
		},
	}
	for _, role := range roles {
		dbQueries.CreateRole(ctx, database.CreateRoleParams{
			Name:        role.Name,
			Description: role.Description,
		})
	}
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		helpers.Logger.Info("DB_URL variable must be defined (use a .env file)", "DB_URL", dbURL)
	}

	sqlDB, err := db.OpenDB(dbURL)
	if err != nil {
		helpers.Logger.Error(err.Error())
		os.Exit(1)
	}
	dbQueries := database.New(sqlDB)

	seedRoles(dbQueries)
}
