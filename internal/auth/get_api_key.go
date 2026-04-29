package auth

import (
	"fmt"
	"os"
)

func GetApiKey() (string, error) {
	key, found := os.LookupEnv("APY_KEY")
	if !found {
		return "", fmt.Errorf("KEY NOT FOUND")
	}
	return key, nil
}
