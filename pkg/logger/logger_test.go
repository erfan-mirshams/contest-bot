package logger

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerTestSuite struct {
	suite.Suite
	buffer *bytes.Buffer
}

func (s *LoggerTestSuite) SetupTest() {
	s.buffer = &bytes.Buffer{}
}

func (s *LoggerTestSuite) TestLogLevels() {
	// Test debug level
	s.Run("Debug level", func() {
		logger := s.setupLoggerWithBuffer("debug")
		logger.Debug("debug message")
		logger.Info("info message")
		s.Contains(s.buffer.String(), "debug message")
		s.Contains(s.buffer.String(), "info message")
	})

	// Test info level
	s.Run("Info level", func() {
		s.buffer.Reset()
		logger := s.setupLoggerWithBuffer("info")
		logger.Debug("debug message")
		logger.Info("info message")
		s.NotContains(s.buffer.String(), "debug message")
		s.Contains(s.buffer.String(), "info message")
	})

	// Test warn level
	s.Run("Warn level", func() {
		s.buffer.Reset()
		logger := s.setupLoggerWithBuffer("warn")
		logger.Info("info message")
		logger.Warn("warn message")
		s.NotContains(s.buffer.String(), "info message")
		s.Contains(s.buffer.String(), "warn message")
	})
}

func (s *LoggerTestSuite) setupLoggerWithBuffer(level string) *slog.Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewTextHandler(s.buffer, opts)
	return slog.New(handler)
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}
