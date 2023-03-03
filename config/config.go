package config

import (
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Telegram Telegram
}

type Telegram struct {
	Token string
}

func NewConfig() *Config {
	c := NewConfigDefaults()
	if err := c.loadConfigFromFile("config.toml"); err != nil {
		log.Println("Unable to load config.toml, loaded defaults...")
	}
	c.applyEnvirontmentVariables()

	return c
}

func NewConfigDefaults() *Config {
	return &Config{}
}

func (c *Config) loadConfigFromFile(path string) error {
	if _, err := toml.DecodeFile(path, c); err != nil {
		return err
	}

	return nil
}

func (c *Config) applyEnvirontmentVariables() {
	applyEnvirontmentVariable("TELEGRAM_TOKEN", &c.Telegram.Token)
}

func applyEnvirontmentVariable(key string, value interface{}) {
	if env, ok := os.LookupEnv(key); ok {
		switch v := value.(type) {
		case *string:
			*v = env
		case *bool:
			if env == "true" || env == "1" {
				*v = true
			} else if env == "false" || env == "0" {
				*v = false
			}
		case *int:
			if number, err := strconv.Atoi(env); err == nil {
				*v = number
			}
		}
	}
}
