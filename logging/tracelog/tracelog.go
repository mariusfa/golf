package tracelog

import (
	"encoding/json"
	"log"
	"time"
)

var tracelogger = NewTraceLogger("")

func SetAppName(appName string) {
	tracelogger.appName = appName
}

func Info(payload string) {
	tracelogger.Info(payload)
}

func Error(payload string) {
	tracelogger.Error(payload)
}

type TraceLogger struct {
	appName string
}

func NewTraceLogger(appName string) *TraceLogger {
	return &TraceLogger{appName: appName}
}

func (tl *TraceLogger) Info(payload string) {
	logLevel := "INFO"
	logType := "TRACE"

	entry := newTraceLog(logLevel, logType, payload, tl.appName)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

func (tl *TraceLogger) Error(payload string) {
	logLevel := "ERROR"
	logType := "TRACE"

	entry := newTraceLog(logLevel, logType, payload, tl.appName)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

type traceLog struct {
	Timestamp string `json:"timestamp"`
	LogLevel  string `json:"log_level"`
	LogType   string `json:"log_type"`
	AppName   string `json:"app_name"`
	Payload   string `json:"payload"`
}

func newTraceLog(logLevel string, logType string, payload string, appName string) *traceLog {
	currentTime := time.Now()
	return &traceLog{
		Timestamp: currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		LogLevel:  logLevel,
		LogType:   logType,
		Payload:   payload,
		AppName:   appName,
	}
}
