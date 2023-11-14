package middlewares

import (
	"fmt"
	"strings"
)

func ValidateAndExtractToken(tokenHeader string) (string, error) {
	if tokenHeader == "" {
		return "", fmt.Errorf("authorization Header is missing")
	}

	accessTokenParts := strings.Split(tokenHeader, " ")

	if len(accessTokenParts) != 2 || accessTokenParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid or missing token")
	}

	return accessTokenParts[1], nil
}
