![Build Status](https://circleci.com/gh/Gamebuildr/gamebuildr-lumberjack.png?circle-token=:circle-token)

# Gamebuildr Lumberjack

### Overview

Lumberjack is the logging system used in all of gamebuildr's microservices.

There are two main ideas behind lumberjack - developers need specification over the format of the log (system logs, user logs, etc.) and how / where the log should be saved.

### Getting Started

To get the lumberjack project simply run ```go get github.com/Gamebuildr/gamebuildr-lumberjack```. This will install the library in your default $GOROOT/src folder.

Code sample for setting up a system logger that saves to file on a local harddrive:

```Go
// Create a new log save system that will persist our log data to a file on the local harddrive
// name the file whatever you like, in this case we've called it system_log_
// File names are automatically appended with the current time and .log
fileLogger := logger.FileLogSave{LogFileName: "system_log_"}

// Create a new logging system to format our data
logger := new(logger.SystemLogger)

// Setup the logsave function to our file logger
logger.LogSave = fileLogger

// There are two methods available for the logging system: Info and Error
logger.Info("Logger System is Saving to File")
logger.Error("Oops, something went wrong")
```

### Papertrail

Papertrail is our remote logging solution - just in case files cannot be persisted (i.e. Docker)

To use papertrail all you need is the URL endpoint you've given on account creation and the App name where the log information is coming from. As above, we'll use the same system logger format and we'll have access to Info() and Error() log output.

Code sample for using papertrail:

```Go
papertrailLog := papertrail.PapertrailLogSave{
	App: "Lumberjack",
	URL: "STUB_ENDPOINT",
}
logs := new(logger.SystemLogger)

logs.LogSave = papertrailLog

logs.Info("Test Saving to Papertrail")
```

Go to the [papertrail go api](http://help.papertrailapp.com/kb/configuration/go-logging/) for more information on using papertrail in Go.