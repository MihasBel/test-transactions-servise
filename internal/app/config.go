package app

import (
	"github.com/MihasBel/test-transactions-servise/adapters/broker"
	"github.com/MihasBel/test-transactions-servise/adapters/pg"
	"github.com/MihasBel/test-transactions-servise/delivery/grpc/server"
	"github.com/MihasBel/test-transactions-servise/delivery/kafka"
)

// Config exported variable to contain config values
var Config Configuration

// Configuration exported type for config
type Configuration struct {
	PG                    pg.Config     `json:"pg"`
	Kafka                 kafka.Config  `json:"kafka"`
	Broker                broker.Config `json:"broker"`
	GRPC                  server.Config `json:"grpc"`
	StartTimeout          int           `json:"start_timeout"`
	StopTimeout           int           `json:"stop_timeout"`
	Address               string        `json:"address"`
	ConsoleLoggingEnabled bool          `json:"console_logging_enabled"`
	FileLoggingEnabled    bool          `json:"file_logging_enabled"`
	LogDirectory          string        `json:"log_directory"`
	LogFilename           string        `json:"log_filename"`
	LogMaxBackups         int           `json:"log_max_backups"`
	LogMaxSize            int           `json:"log_max_size"`
	LogMaxAge             int           `json:"log_max_age"`
}
