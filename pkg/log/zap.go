package log

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	Error error
}

func New() *Logger {
	mode, ok := os.LookupEnv("LOG_MODE")
	if !ok {
		mode = "dev"
		fmt.Println("LOG_MODE not set, defaulting to DEV")
	}

	l := &Logger{}
	switch strings.ToLower(mode) {
	case "dev":
		l.NewDevLogger()
	case "prod":
		l.NewProdLogger()
	default:
		l.NewProdLogger()
	}

	return l
}

func (l *Logger) NewDevLogger() error {
	err := l.NewDevelopmentColorizedLogger()
	if err != nil {
		l.Error = err
		return err
	}
	return nil
}

func (l *Logger) NewProdLogger() error {
	err := l.NewProductionColorizedLogger()
	if err != nil {
		l.Error = err
		return err
	}
	return nil
}

func (l *Logger) NewDevelopmentColorizedLogger() error {
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		EncoderConfig:    cfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		l.Error = err
		return err
	}

	l.Logger = logger
	return nil
}

func (l *Logger) NewDevelopmentJsonLogger() error {
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		EncoderConfig:    cfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		l.Error = err
		return err
	}

	l.Logger = logger
	return nil
}

func (l *Logger) NewProductionColorizedLogger() error {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		EncoderConfig:    cfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		l.Error = err
		return err
	}

	l.Logger = logger
	return nil
}

func (l *Logger) NewProductionJsonLogger() error {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		EncoderConfig:    cfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		l.Error = err
		return err
	}

	l.Logger = logger
	return nil
}
