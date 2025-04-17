package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gotd/td/telegram"
)

func GetTGFileSessionStorage(appID int) (*telegram.FileSessionStorage, error) {
	dirName := strconv.Itoa(appID)
	dirPath := filepath.Join(".sessions", dirName)

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return nil, fmt.Errorf("os.MkdirAll: %w", err)
	}

	return &telegram.FileSessionStorage{
		Path: filepath.Join(dirPath, "session.json"),
	}, nil
}
