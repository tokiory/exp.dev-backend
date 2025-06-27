package config

import (
	"log/slog"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type databaseConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     uint16
}

type config struct {
	Database databaseConfig
}

func NewConfig(log *slog.Logger, path string) *config {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		panic("Error parsing DB_PORT: " + err.Error())
	}

	conf := &config{
		Database: databaseConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     uint16(port),
			Name:     os.Getenv("DB_NAME"),
		},
	}

	if os.Getenv("DEBUG") == "1" {
		log.Debug("Config loaded",
			slog.Group("database",
				slog.String("user", conf.Database.User),
				slog.String("host", conf.Database.Host),
				slog.Uint64("port", uint64(conf.Database.Port)),
				slog.String("name", conf.Database.Name),
				slog.String("password", conf.Database.Password),
			),
		)
	}
	return conf
}
