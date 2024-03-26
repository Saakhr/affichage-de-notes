package main

import (
	"database/sql"
	"log/slog"
	"os"
)

type ApiConf struct {
	db *sql.DB
}

func main() {

	db := setupDB()
	defer db.Close()
	ApiCFG := ApiConf{
		db: db,
	}
	// Run your server.
	if err := runServer(ApiCFG); err != nil {
		slog.Error("Failed to start server!", "details", err.Error())
		os.Exit(1)
	}
}
