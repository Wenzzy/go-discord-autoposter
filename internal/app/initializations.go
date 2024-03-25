package app

import (
	"context"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wenzzy/go-discord-autoposter/internal/app/di"
	"github.com/wenzzy/go-discord-autoposter/internal/models"
	"github.com/wenzzy/go-discord-autoposter/internal/utils/logger"
)

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		// a.initConfig,
		a.initServiceProvider,
		a.initLogger,
		a.initAutoPoster,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Currently not used
// func (a *App) initConfig(_ context.Context) error {
// 	err := config.Load(".env")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (a *App) initServiceProvider(_ context.Context) error {
	a.diContainer = di.NewContainer()
	return nil
}

func (a *App) initAutoPoster(ctx context.Context) error {
	c := cron.New(cron.WithSeconds())

	messages := a.diContainer.MessagesConfig().Messages()
	// TODO: Create AutoPoster struct and separate the logic
	_, err := c.AddFunc("*/10 * * * * *", func() {

		defer func() {
			if r := recover(); r != nil {
				logger.Error("Panic recovered in sendMessagesToDiscordChannel, %v", zap.Any("r", r))
			}
		}()

		for _, msg := range messages {
			err := a.diContainer.MessageService(ctx).Send(ctx, models.SendMessage{
				ChannelID:   msg.ChannelID,
				Content:     &msg.Text,
				Interval:    msg.Interval,
				Attachments: msg.Files,
				Topic:       msg.Topic,
				ChannelName: msg.ChannelName,
			})
			if err != nil {
				logger.Error("Failed to send message", zap.Error(err))
			}
		}
	})

	if err != nil {
		return err
	}

	a.autoPoster = c
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	var zapLevel zapcore.Level
	if err := zapLevel.Set(a.diContainer.LoggerConfig().Level()); err != nil {
		return err
	}

	level := zap.NewAtomicLevelAt(zapLevel)

	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   a.diContainer.LoggerConfig().Path(),
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(cfg)
	out := stdout

	if a.diContainer.AppConfig().Env() == "prod" {
		cfg = zap.NewProductionEncoderConfig()
		cfg.TimeKey = "timestamp"
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder = zapcore.NewJSONEncoder(cfg)
		out = file
	}

	core := zapcore.NewCore(encoder, out, level)

	// Sampling
	samplingCore := zapcore.NewSamplerWithOptions(
		core,
		time.Second, // interval
		3,           // log first 3 entries
		1,           // thereafter log zero entires within the interval
	)

	logger.Init(samplingCore)

	return nil
}
