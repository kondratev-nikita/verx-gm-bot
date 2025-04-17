package utils

import (
	"strings"
)

func BytesToStrSlice(v []byte) []string {
	return strings.Split(string(v), "\n\n")
}
