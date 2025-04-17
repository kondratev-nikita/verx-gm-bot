package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type TGAcc struct {
	ApiID          int
	ApiHash        string
	TargetUsername string
}

type Config struct {
	TGAcc TGAcc
}

func New() (Config, error) {
	apiIDStr := os.Getenv("TG_ACC_API_ID")
	if apiIDStr == "" {
		return Config{}, errors.New("TG_ACC_API_ID not set")
	}
	apiIDInt, err := strconv.Atoi(apiIDStr)
	if err != nil {
		return Config{}, fmt.Errorf("TG_ACC_API_ID is not a number: %w", err)
	}

	apiHash := os.Getenv("TG_ACC_API_HASH")
	if apiHash == "" {
		return Config{}, errors.New("TG_ACC_API_HASH not set")
	}

	targetUsername := os.Getenv("TG_ACC_TARGET_USERNAME")
	if targetUsername == "" {
		return Config{}, errors.New("TG_ACC_TARGET_USERNAME not set")
	}

	return Config{
		TGAcc: TGAcc{
			ApiID:          apiIDInt,
			ApiHash:        apiHash,
			TargetUsername: targetUsername,
		},
	}, nil
}
