package transactionlog

import (
	"encoding/json"
	"log"
	"time"
)

var TransLog TransLogger

type TransLogger struct {
	appName string
}

func NewTransLogger(appName string) TransLogger {
	log.SetFlags(0)
	return TransLogger{appName: appName}
}

func (al *TransLogger) Info(payload string, requestMethod string, event string, status string, durationMs string) {
	logLevel := "INFO"
	logType := "transaction"

	entry := newTransLog(logLevel, logType, payload, al.appName, requestMethod, event, status, durationMs)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

type transLog struct {
	Timestamp     string `json:"timestamp"`
	LogLevel      string `json:"log_level"`
	LogType       string `json:"log_type"`
	AppName       string `json:"app_name"`
	RequestMethod string `json:"request_method"`
	Event         string `json:"event"`
	Status        string `json:"status"`
	Payload       string `json:"payload"`
	DurationMs    string `json:"duration_ms"`
}

func newTransLog(logLevel string, logType string, payload string, appName string, requestMethod string, event string, status string, durationMs string) *transLog {
	currentTime := time.Now()
	return &transLog{
		Timestamp:     currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		LogLevel:      logLevel,
		LogType:       logType,
		Payload:       payload,
		AppName:       appName,
		RequestMethod: requestMethod,
		Event:         event,
		Status:        status,
		DurationMs:    durationMs,
	}
}
