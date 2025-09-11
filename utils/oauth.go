package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateUsername(base string) string {
	username := strings.ToLower(strings.ReplaceAll(base, " ", ""))
	if username == "" {
		username = "user"
	}
	return fmt.Sprintf("%s_%s", username, uuid.New().String()[:8])
}
