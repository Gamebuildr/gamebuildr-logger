package papertrail

import "log/syslog"
import "fmt"

type PapertrailLogSave struct {
	App string
	URL string
}

func (logsave PapertrailLogSave) SaveLogData(data string) {
	w, err := syslog.Dial("udp", logsave.URL, syslog.LOG_EMERG|syslog.LOG_KERN, logsave.App)
	if err != nil {
		fmt.Printf("Papertrail Save Error: %v", err.Error())
		w.Err(err.Error())
	}
	w.Info(data)
}
