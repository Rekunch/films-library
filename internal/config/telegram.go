package config

import (
	"errors"
	"strconv"
)

const (
	telegramTokenEnvName   = "TELEGRAM_TOKEN"
	telegramTimeoutEnvName = "TELEGRAM_TIMEOUT"

	defaultTelegramTimeout = 60
)

// TelegramBotConfig ...
type TelegramBotConfig interface {
	TelegramToken() string
	Timeout() int
}

type telegramConfig struct {
	telegramToken string
	timeout       int
}

// GetTelegramBotConfig ...
func GetTelegramBotConfig(isStgEnv bool) (TelegramBotConfig, error) {
	token := get(telegramTokenEnvName)
	if len(token) == 0 {
		return nil, errors.New("telegram token not found")
	}

	timeoutStr := get(telegramTimeoutEnvName)
	timeout, err := strconv.ParseInt(timeoutStr, 10, 64)
	if err != nil || timeout == 0 {
		timeout = defaultTelegramTimeout
	}

	return &telegramConfig{
		telegramToken: token,
		timeout:       int(timeout),
	}, nil
}

// TelegramToken ...
func (c *telegramConfig) TelegramToken() string {
	return c.telegramToken
}

// Timeout ...
func (c *telegramConfig) Timeout() int {
	return c.timeout
}
