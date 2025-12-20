package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/zexk/mercurylampe/internal/cat"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(os.Getenv("TG_TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/cat", bot.MatchTypeExact, catHandler)

	b.Start(ctx)
}

func catHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := cat.GetCat()

	params := &bot.SendPhotoParams{
		ChatID: update.Message.Chat.ID,
		Photo:  &models.InputFileUpload{Filename: "cat.png", Data: bytes.NewReader(data)},
	}

	b.SendPhoto(ctx, params)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		fmt.Printf("%s:\n %s\n", update.Message.Chat.Username, update.Message.Text)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}
}
