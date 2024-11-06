package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
	return &Config{
		Port: os.Getenv("PORT"),
	}
}
