package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"tenders/pkg/api"
	"tenders/pkg/env"
	"tenders/pkg/handler"
	"tenders/pkg/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	pmigrate "github.com/peterldowns/pgmigrate"
)

func setupDb(pgUrl string) *pgxpool.Pool {
	migrPath := "./migrations"

	db, err := pgxpool.New(context.Background(), pgUrl)
	if err != nil {
		log.Panicln(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		log.Panicln(err)
	}
	log.Println("connected to the db")

	sqlDb, err := sql.Open("pgx", pgUrl)
	if err != nil {
		log.Panicln(err)
	}
	defer sqlDb.Close()

	migrations, err := pmigrate.Load(os.DirFS(migrPath))
	if err != nil {
		log.Panicln(err)
	}
	migrator := pmigrate.NewMigrator(migrations)
	_, err = migrator.Migrate(context.Background(), sqlDb)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("migration successful")

	return db
}

func main() {
	env := env.Read()

	db := setupDb(env.PostgresUrl)
	defer db.Close()

	handler := handler.New(db)
	server, err := api.NewServer(handler, api.WithMiddleware(middleware.Logging()))
	if err != nil {
		log.Panicln(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", server))

	log.Println("starting the server")
	err = http.ListenAndServe(env.ServerUrl, mux)
	if err != nil {
		log.Panicln(err)
	}
}
