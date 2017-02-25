package logger

import "testing"

func TestInterfaceCanLogInfo(t *testing.T) {
	mockLog := MockLog{Test: t}
	message := "mock message"
	info := mockLog.Info(message)

	if info != message {
		t.Errorf("Expected %v, but got %v", message, info)
	}
}

func TestInterfaceCanLogErrors(t *testing.T) {
	mockLog := MockLog{Test: t}
	errMessage := "mock error"
	err := mockLog.Error(errMessage)

	if err != errMessage {
		t.Errorf("Expected %v, but got %v", errMessage, err)
	}
}
