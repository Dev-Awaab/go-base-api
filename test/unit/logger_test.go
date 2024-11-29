package logger_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/Dev-Awaab/go-base-api/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestLoggerInitialization(t *testing.T) {
	log, err := logger.InitLogger(logger.InfoLevel, "stdout")
	require.NoError(t, err, "Logger initialization should not produce an error")
	require.NotNil(t, log, "Logger instance should not be nil")
}

func TestLoggerInfo(t *testing.T) {
	var buf bytes.Buffer
	log, err := logger.InitLogger(logger.InfoLevel, "stdout")
	require.NoError(t, err, "Logger initialization failed")
	log.SetOutput(&buf)

	log.Info("Test info log message")

	require.Contains(t, buf.String(), "[INFO] Test info log message", "Log output should contain the expected info message")
}

func TestLoggerError(t *testing.T) {
	var buf bytes.Buffer
	log, err := logger.InitLogger(logger.ErrorLevel, "stdout")
	require.NoError(t, err, "Logger initialization failed")
	log.SetOutput(&buf)

	log.Error("Test error log message")

	require.Contains(t, buf.String(), "[ERROR] Test error log message", "Log output should contain the expected error message")
}

func TestLoggerToFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "logfile-*.log")
	require.NoError(t, err, "Failed to create temporary log file")
	defer os.Remove(tmpfile.Name())

	log, err := logger.InitLogger(logger.InfoLevel, tmpfile.Name())
	require.NoError(t, err, "Logger initialization failed")

	log.Info("Test log to file")

	err = tmpfile.Sync()
	require.NoError(t, err, "Failed to sync log file")

	content, err := os.ReadFile(tmpfile.Name())
	require.NoError(t, err, "Failed to read log file")

	require.Contains(t, string(content), "[INFO] Test log to file", "Log file should contain the expected info message")
}

func TestLoggerLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	log, err := logger.InitLogger(logger.WarnLevel, "stdout")
	require.NoError(t, err, "Logger initialization failed")
	log.SetOutput(&buf)

	log.Debug("This is a debug message")
	log.Info("This is an info message")
	log.Warn("This is a warning message")

	require.NotContains(t, buf.String(), "debug", "Debug log should not appear at WarnLevel")
	require.NotContains(t, buf.String(), "info", "Info log should not appear at WarnLevel")
	require.Contains(t, buf.String(), "This is a warning message", "Warning message should appear at WarnLevel")
}