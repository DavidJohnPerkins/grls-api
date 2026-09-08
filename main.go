package main

import (
	"context"
	"dperkins/grls-api/api"
	"dperkins/grls-api/config"
	"dperkins/grls-api/store"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dbx, err := store.InitSharedDB(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	grlsStore := store.NewSqlServerGrlsStore(dbx)
	server := api.NewServer(cfg.HTTPServer, grlsStore)
	server.Start(ctx)
}
