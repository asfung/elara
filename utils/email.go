package utils

import (
	"fmt"
	"strings"
)

func GetEmailPrefix(email string) (string, error) {
	if !strings.Contains(email, "@") {
		return "", fmt.Errorf("invalid email: missing @")
	}
	parts := strings.SplitN(email, "@", 2)
	if parts[0] == "" {
		return "", fmt.Errorf("invalid email: missing prefix")
	}
	return parts[0], nil
}
