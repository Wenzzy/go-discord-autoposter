package main

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/wenzzy/go-discord-autoposter/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
