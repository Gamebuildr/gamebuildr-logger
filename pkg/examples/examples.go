package examples

import "github.com/Gamebuildr/gamebuildr-lumberjack/pkg/logger"

// LoggerExample shows how to implement the logger interface
func LoggerExample() {
	// Create a new log save system that will persist our log data
	fileLogger := logger.FileLogSave{LogFileName: "system_log_"}

	// Create a new logging system to format our data
	logger := new(logger.SystemLogger)

	// Setup the logsave function to our file logger
	logger.LogSave = fileLogger

	// Use of the logger will store data into a file
	logger.Info("Logger System is Saving to File")
}
