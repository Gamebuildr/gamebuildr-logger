package logger

import "testing"

// MockLog is a logger for use inside tests
type MockLog struct {
	Test *testing.T
}

// Info prints data to console when logging in tests
func (log MockLog) Info(message string) string {
	log.Test.Logf(message)
	return message
}

// Error prints errors to console when logging in tests
func (log MockLog) Error(err string) string {
	log.Test.Logf("Error in test: %v", err)
	return err
}
