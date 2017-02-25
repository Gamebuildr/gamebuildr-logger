package logger

import (
	"testing"
	"time"
)

func TestSystemInfoLoggerReturnsCorrectInfo(t *testing.T) {
	log := new(SystemLogger)
	mockInfo := log.Info("stub message")
	time := time.Now().Local()
	timeString := time.Format("Mon Jan _2 15:04:05 UTC 2006")
	testInfo := SystemLogTag + " " + timeString + ": stub message"
	if mockInfo != testInfo {
		t.Errorf("Expected: %v, got: %v", testInfo, mockInfo)
	}
}

func TestSystemErrorLoggerReturnsCorrectError(t *testing.T) {
	log := new(SystemLogger)
	mockError := log.Error("stub error")
	time := time.Now().Local()
	timeString := time.Format("Mon Jan _2 15:04:05 UTC 2006")
	testError := SystemErrorTag + " " + timeString + ": stub error"

	if mockError != testError {
		t.Errorf("Expected: %v, got: %v", mockError, testError)
	}
}
