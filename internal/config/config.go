package config

import (
	"github.com/mohammadne/phone-book/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
