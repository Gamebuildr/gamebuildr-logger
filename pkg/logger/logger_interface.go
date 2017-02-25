package logger

// LogSave is the main system to specify
// how to save log information
type LogSave interface {
	SaveLogData(message string)
}

// Logger is the main interface struct
// for different types of loggin
type Logger struct {
	LogSave LogSave
}
