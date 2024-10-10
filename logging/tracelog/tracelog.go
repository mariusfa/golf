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

func Info(payload string, userId string, requestId string) {
	tracelogger.Info(payload, userId, requestId)
}

func Error(payload string, userId string, requestId string) {
	tracelogger.Error(payload, userId, requestId)
}

type TraceLogger struct {
	appName string
}

func NewTraceLogger(appName string) *TraceLogger {
	return &TraceLogger{appName: appName}
}

func GetLogger() *TraceLogger {
	return tracelogger
}

func (tl *TraceLogger) Info(payload string, userId string, requestId string) {
	logLevel := "INFO"
	logType := "TRACE"

	entry := newTraceLog(logLevel, logType, payload, tl.appName, userId, requestId)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

func (tl *TraceLogger) Error(payload string, userId string, requestId string) {
	logLevel := "ERROR"
	logType := "TRACE"

	entry := newTraceLog(logLevel, logType, payload, tl.appName, userId, requestId)
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
	UserId    string `json:"user_id"`
	RequestId string `json:"request_id"`
}

func newTraceLog(logLevel string, logType string, payload string, appName string, userId string, requestId string) *traceLog {
	currentTime := time.Now()
	return &traceLog{
		Timestamp: currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		LogLevel:  logLevel,
		LogType:   logType,
		Payload:   payload,
		AppName:   appName,
		UserId:    userId,
		RequestId: requestId,
	}
}
