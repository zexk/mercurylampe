package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/zexk/mercurylampe/internal/cat"
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}
}

func catHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := cat.GetCat()

	params := &bot.SendPhotoParams{
		ChatID: update.Message.Chat.ID,
		Photo:  &models.InputFileUpload{Data: data},
	}

	b.SendPhoto(ctx, params)
}

func gcatHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := cat.GetGifCat()

	params := &bot.SendAnimationParams{
		ChatID:    update.Message.Chat.ID,
		Animation: &models.InputFileUpload{Filename: "file.gif", Data: data},
	}

	b.SendAnimation(ctx, params)
}

func tcatHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "what should the cat say?",
	})
	data := cat.GetTextCat(update.Message.Text)

	params := &bot.SendPhotoParams{
		ChatID: update.Message.Chat.ID,
		Photo:  &models.InputFileUpload{Data: data},
	}

	b.SendPhoto(ctx, params)
}
