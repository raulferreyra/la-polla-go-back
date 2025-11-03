package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv    string
	HttpAddr  string
	DBURL     string
	JWTSecret string
	AESKey    string // 32 bytes base64/hex
}

func Load() *Config {
	viper.SetDefault("APP_ENV", "dev")
	viper.SetDefault("HTTP_ADDR", ":8080")
	viper.AutomaticEnv()

	cfg := &Config{
		AppEnv:    viper.GetString("APP_ENV"),
		HttpAddr:  viper.GetString("HTTP_ADDR"),
		DBURL:     viper.GetString("DATABASE_URL"),
		JWTSecret: viper.GetString("JWT_SECRET"),
		AESKey:    viper.GetString("AES_KEY"),
	}
	if cfg.DBURL == "" || cfg.JWTSecret == "" || len(cfg.AESKey) < 32 {
		log.Fatal("Env incompleto: DATABASE_URL, JWT_SECRET, AES_KEY(32+)")
	}
	_ = time.Now() // evita lint
	return cfg
}
