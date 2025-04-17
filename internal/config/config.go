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
	apiIdStr := os.Getenv("TG_ACC_API_ID")
	if apiIdStr == "" {
		return Config{}, errors.New("TG_ACC_API_ID not set")
	}
	apiIdInt, err := strconv.Atoi(apiIdStr)
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
			ApiID:          apiIdInt,
			ApiHash:        apiHash,
			TargetUsername: targetUsername,
		},
	}, nil
}
