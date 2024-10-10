package middlewarelog

import (
	"encoding/json"
	"log"
	"time"
)

var middlewareLogger = NewMiddlewareLogger("")

func SetAppName(appName string) {
	middlewareLogger.appName = appName
}

func Info(payload string, requestId string) {
	middlewareLogger.Info(payload, requestId)
}

func Error(payload string, requestId string) {
	middlewareLogger.Error(payload, requestId)
}

type MiddlewareLogger struct {
	appName string
}

func NewMiddlewareLogger(appName string) *MiddlewareLogger {
	log.SetFlags(0)
	return &MiddlewareLogger{appName: appName}
}

func GetLogger() *MiddlewareLogger {
	return middlewareLogger
}

func (al *MiddlewareLogger) Info(payload string, requestId string) {
	logLevel := "INFO"
	logType := "MIDDLWARE"

	entry := newMiddlewareLog(logLevel, logType, payload, requestId, al.appName)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

type middlewareLog struct {
	Timestamp string `json:"timestamp"`
	LogLevel  string `json:"log_level"`
	LogType   string `json:"log_type"`
	AppName   string `json:"app_name"`
	Payload   string `json:"payload"`
	RequestId string `json:"request_id"`
}

func (al *MiddlewareLogger) Error(payload string, requestId string) {
	logLevel := "ERROR"
	logType := "MIDDLEWARE"

	entry := newMiddlewareLog(logLevel, logType, payload, requestId, al.appName)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

func newMiddlewareLog(logLevel string, logType string, payload string, requestId string, appName string) *middlewareLog {
	currentTime := time.Now()
	return &middlewareLog{
		Timestamp: currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		LogLevel:  logLevel,
		LogType:   logType,
		Payload:   payload,
		AppName:   appName,
		RequestId: requestId,
	}
}
