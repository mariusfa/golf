package accesslog

import (
	"encoding/json"
	"log"
	"time"
)

var AccessLog AccessLogger

type AccessLogger struct {
	appName string
}

func NewAccessLogger(appName string) AccessLogger {
	return AccessLogger{appName: appName}
}

func (al *AccessLogger) Info(durationMs int, status int, requestPath string, requestMethod string) {
	logLevel := "INFO"
	logType := "ACCESS"

	entry := newAccessLog(logLevel, logType, durationMs, status, requestPath, requestMethod, al.appName)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

type accessLog struct {
	Timestamp     string `json:"timestamp"`
	DurationMs    int    `json:"duration_ms"`
	LogLevel      string `json:"log_level"`
	LogType       string `json:"log_type"`
	AppName       string `json:"app_name"`
	Status        int    `json:"status"`
	RequestPath   string `json:"request_path"`
	RequestMethod string `json:"request_method"`
}

func newAccessLog(logLevel string, logType string, durationMs int, status int, requestPath string, requestMethod string, appName string) *accessLog {
	currentTime := time.Now()
	return &accessLog{
		Timestamp:     currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		DurationMs:    durationMs,
		LogLevel:      logLevel,
		LogType:       logType,
		Status:        status,
		RequestPath:   requestPath,
		RequestMethod: requestMethod,
		AppName:       appName,
	}
}
