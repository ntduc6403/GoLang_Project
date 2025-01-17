package config

import "github.com/caarlos0/env/v6"

type Config struct {
	HTTPServer HTTPServerConfig
	Logger     LoggerConfig
	Database   DatabaseConfig
	PG         PGConfig
}

type PGConfig struct {
	URL string `env:"PG_URL" envDefault:"postgres://postgres:postgres@127.0.0.1:5454/music_db?sslmode=disable"`
}

type DatabaseConfig struct {
	Driver string `env:"DB_DRIVER" envDefault:"postgres"`
}

type HTTPServerConfig struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Mode string `env:"MODE" envDefault:"development"`
}

type LoggerConfig struct {
	Level    string `env:"LOGGER_LEVEL" envDefault:"debug"`
	Mode     string `env:"MODE" envDefault:"development"`
	Encoding string `env:"LOGGER_ENCODING" envDefault:"console"`
}

// Load loads the configuration from the environment variables.
func Load() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
