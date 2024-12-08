package config

import (
	"github.com/yay99857/mineduo-server/pkg/database"
)

func Init() {
	database.ConnectDatabase()
}

// Get new log instace
func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
