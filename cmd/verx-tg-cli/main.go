package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"github.com/kondratev-nikita/verx-gm-bot/internal/utils"
	"golang.org/x/term"
)

func main() {
	phone := flag.String("p", "", "Номер телефона тг аккаунта (79112223434)")
	apiID := flag.Int("id", 0, "API ID тг аккаунта")
	apiHash := flag.String("h", "", "API Hash тг аккаунта")
	usePass := flag.Bool("pw", false, "Вход с паролем, по умолчанию без 2FA")

	flag.Parse()

	if phone == nil || *phone == "" {
		fatal("Номер обязателен")
	}
	if apiID == nil || *apiID == 0 {
		fatal("API ID обязателен")
	}
	if apiHash == nil || *apiHash == "" {
		fatal("API Hash обязателен")
	}

	ss, err := utils.GetTGFileSessionStorage(*apiID)
	if err != nil {
		fatal(fmt.Errorf("не вышло создать хранилище сессий: %w", err).Error())
	}

	tc := telegram.NewClient(*apiID, *apiHash, telegram.Options{
		SessionStorage: ss,
	})

	ctx := context.Background()

	if *usePass {
		fmt.Print("Введите пароль: ")
		passB, errR := term.ReadPassword(syscall.Stdin)
		if errR != nil {
			fatal(errR.Error())
		}
		fmt.Println()

		pass := strings.TrimSpace(string(passB))
		err = loginWithPass(ctx, tc, *phone, pass)
	} else {
		err = login(ctx, tc, *phone)
	}

	if err != nil {
		fatal(fmt.Errorf("не вышло войти: %w", err).Error())
	}
}

func fatal(v string) {
	fmt.Println("Error:", v)
	flag.Usage()
	os.Exit(1)
}

func loginWithPass(ctx context.Context, tc *telegram.Client, phone, pass string) error {
	codeAsk := func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
		fmt.Print("Введите код с TG: ")
		code, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return "", err
		}
		code = strings.ReplaceAll(code, "\n", "")
		return code, nil
	}

	return tc.Run(ctx, func(ctx context.Context) error {
		return auth.NewFlow(
			auth.Constant(phone, pass, auth.CodeAuthenticatorFunc(codeAsk)),
			auth.SendCodeOptions{},
		).Run(ctx, tc.Auth())
	})
}

func login(ctx context.Context, tc *telegram.Client, phone string) error {
	codeAsk := func(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
		fmt.Print("Введите код с TG: ")
		code, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return "", err
		}
		code = strings.ReplaceAll(code, "\n", "")
		return code, nil
	}

	return tc.Run(ctx, func(ctx context.Context) error {
		return auth.NewFlow(
			auth.CodeOnly(phone, auth.CodeAuthenticatorFunc(codeAsk)),
			auth.SendCodeOptions{},
		).Run(ctx, tc.Auth())
	})
}
