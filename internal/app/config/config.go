package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
}

func Init() (*Config, error) {
	cfg := new(Config)
	err := parseEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseEnv(cfg *Config) error {
	//err := viper.BindEnv("telegram_token")
	//if err != nil {
	//	return err
	//}
	//cfg.TelegramToken = viper.GetString("telegram_token")
	godotenv.Load()
	cfg.TelegramToken = os.Getenv("telegram_token")

	return nil
}
