package app

import (
	"context"
	"log"
	"sync"

	"github.com/robfig/cron/v3"

	"github.com/wenzzy/go-discord-autoposter/internal/app/di"
	"github.com/wenzzy/go-discord-autoposter/internal/utils/closer"
)

type App struct {
	diContainer *di.Container
	autoPoster  *cron.Cron
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		err := a.runAutoPoster()
		if err != nil {
			log.Fatalf("failed to run AutoPoster: %v", err)
		}
	}()

	wg.Wait()

	return nil
}
