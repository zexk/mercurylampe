package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func botInit() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
		bot.WithMiddlewares(logMessage),
	}

	b, err := bot.New(os.Getenv("TG_TOKEN"), opts...)
	if nil != err {
		log.Fatal(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/cat", bot.MatchTypeExact, catHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/gcat", bot.MatchTypeExact, gcatHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/tcat", bot.MatchTypeExact, tcatHandler)

	b.Start(ctx)
}

func logMessage(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("{%d} [%s]: %s", update.Message.From.ID, update.Message.From.Username, update.Message.Text)
		}
		next(ctx, b, update)
	}
}
