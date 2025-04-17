package main

import (
	_ "embed"
	"github.com/kondratev-nikita/verx-gm-bot/internal/services/msg_gen"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/joho/godotenv"
	"github.com/kondratev-nikita/verx-gm-bot/internal/bot"
	"github.com/kondratev-nikita/verx-gm-bot/internal/config"
	"github.com/kondratev-nikita/verx-gm-bot/internal/utils"

	"github.com/go-co-op/gocron/v2"
)

var (
	//go:embed assets/gm_variants.txt
	gmMsgs []byte
	//go:embed assets/gn_variants.txt
	gnMsgs []byte
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	cfg, err := config.New()
	if err != nil {
		l.Error("failed get config", slog.String("err", err.Error()))
		os.Exit(1)
	}

	mskLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		l.Error("failed to load MSK timezone", slog.String("err", err.Error()))
		os.Exit(1)
	}

	scheduler, err := gocron.NewScheduler(gocron.WithLocation(mskLocation))
	if err != nil {
		l.Error("failed get scheduler", slog.String("err", err.Error()))
		os.Exit(1)
	}
	msgGenSvc, err := msg_gen.New(msg_gen.Config{
		MessagesCfg: msg_gen.MessagesConfig{
			GoodNight:   utils.BytesToStrSlice(gnMsgs),
			GoodMorning: utils.BytesToStrSlice(gmMsgs),
		},
	})
	if err != nil {
		l.Error("failed get msg_gen service", slog.String("err", err.Error()))
		os.Exit(1)
	}

	ss, err := utils.GetTGFileSessionStorage(cfg.TGAcc.ApiID)
	if err != nil {
		l.Error("failed get tg session storage", slog.String("err", err.Error()))
		os.Exit(1)
	}
	tc := telegram.NewClient(cfg.TGAcc.ApiID, cfg.TGAcc.ApiHash, telegram.Options{
		SessionStorage: ss,
	})

	b := bot.New(msgGenSvc, tc, cfg.TGAcc.TargetUsername, l)
	err = b.Startup(scheduler)
	if err != nil {
		l.Error("failed bot startup", slog.String("err", err.Error()))
		os.Exit(1)
	}

	scheduler.Start()
	l.Info("verx-gm-bot started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	l.Info("shutting down")
}

func init() {
	_ = godotenv.Load()
}
