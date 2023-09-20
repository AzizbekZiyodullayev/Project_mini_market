package main

import (
	"context"
	"file/api"
	"file/api/handler"
	"file/config"

	"file/pkg/logger"
	"file/storage/db"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := db.NewStorage(context.Background(), cfg)
	if err != nil {
		return
	}
	h := handler.NewHandler(strg, log)

	r := api.NewServer(h)
	r.Run(":8080")
}
