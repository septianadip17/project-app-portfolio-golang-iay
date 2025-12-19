package logger

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	// Menggunakan config production agar formatnya JSON dan rapi
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"} // Output ke terminal

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
