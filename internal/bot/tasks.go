package bot

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/gotd/td/telegram/message"
	"log/slog"
)

func (b *Bot) taskSendGoodMorningMsg() gocron.Task {
	return gocron.NewTask(
		func(ctx context.Context) {
			b.l.Info("start sending good morning message")

			err := b.sendMessageToTargetUser(ctx, b.msgGenSvc.GenGoodMorning())
			if err != nil {
				b.l.Error("failed to send good morning message", slog.String("err", err.Error()))
			} else {
				b.l.Info("sent good morning message")
			}
		},
		context.Background(),
	)
}

func (b *Bot) taskSendGoodNightMsg() gocron.Task {
	return gocron.NewTask(
		func(ctx context.Context) {
			b.l.Info("start sending good night message")

			err := b.sendMessageToTargetUser(ctx, b.msgGenSvc.GenGoodNight())
			if err != nil {
				b.l.Error("failed to send good night message", slog.String("err", err.Error()))
			} else {
				b.l.Info("sent good night message")
			}
		},
		context.Background(),
	)
}

func (b *Bot) sendMessageToTargetUser(ctx context.Context, msg string) error {
	err := b.tc.Run(ctx, func(ctx context.Context) error {
		sender := message.NewSender(b.tc.API())

		_, err := sender.Resolve(b.tgAccTargetUsername).Text(ctx, msg)
		if err != nil {
			return fmt.Errorf("failed to send msg: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("sendMessageToTargetUser: %w", err)
	}

	return nil
}
