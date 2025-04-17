package bot

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"log/slog"

	"github.com/gotd/td/telegram"
	"github.com/kondratev-nikita/verx-gm-bot/internal/msg_gen"
)

type Bot struct {
	msgGenSvc           msg_gen.Service
	tc                  *telegram.Client
	tgAccTargetUsername string
	l                   *slog.Logger
}

func New(
	msgGenSvc msg_gen.Service,
	tc *telegram.Client,
	tgAccTargetUsername string,
	l *slog.Logger,
) *Bot {
	return &Bot{
		msgGenSvc:           msgGenSvc,
		tc:                  tc,
		tgAccTargetUsername: tgAccTargetUsername,
		l:                   l,
	}
}

func (b *Bot) Startup(s gocron.Scheduler) error {
	_, err := s.NewJob(
		gocron.CronJob("0 23 * * *", true), // at 23:00 every day
		b.taskSendGoodNightMsg(),
	)
	if err != nil {
		return fmt.Errorf("job.taskSendGoodNightMsg: %w", err)
	}

	_, err = s.NewJob(
		gocron.CronJob("0 10 * * *", true), // at 10:00 every day
		b.taskSendGoodMorningMsg(),
	)
	if err != nil {
		return fmt.Errorf("job.taskSendGoodMorningMsg: %w", err)
	}

	return nil
}
