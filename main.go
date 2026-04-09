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

	store := store.NewSqlServerGrlsStore(cfg.DatabaseURL)
	server := api.NewServer(cfg.HTTPServer, store)
	server.Start(ctx)
}
