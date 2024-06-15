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

func NewTransLogger(appName string) *TransLogger {
	log.SetFlags(0)
	return &TransLogger{appName: appName}
}

func (al *TransLogger) RequestInfo(requestId string, requestMethod string, requestPath string, requestBody string) {
	logLevel := "INFO"
	logType := "transaction"

	entry := newRequestLog(logLevel, logType, al.appName, requestMethod, requestPath, requestId, requestBody)
	entry.RequestBody = requestBody
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

func (al *TransLogger) ResponseInfo(requestId string, durationMs string, status int, responseBody string) {
	logLevel := "INFO"
	logType := "transaction"

	entry := newResponseLog(logLevel, logType, responseBody, al.appName, status, durationMs, requestId)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

type requestLog struct {
	Timestamp     string `json:"timestamp"`
	LogLevel      string `json:"log_level"`
	LogType       string `json:"log_type"`
	Event         string `json:"event"`
	AppName       string `json:"app_name"`
	RequestMethod string `json:"request_method"`
	RequestPath   string `json:"request_path"`
	RequestId     string `json:"request_id"`
	RequestBody   string `json:"request_body"`
}

func newRequestLog(logLevel string, logType string, appName string, requestMethod string, requestPath string, requestId string, requestBody string) *requestLog {
	currentTime := time.Now()
	return &requestLog{
		Timestamp:     currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		LogLevel:      logLevel,
		LogType:       logType,
		Event:         "request",
		AppName:       appName,
		RequestMethod: requestMethod,
		RequestPath:   requestPath,
		RequestId:     requestId,
		RequestBody:   requestBody,
	}
}

type responseLog struct {
	Timestamp    string `json:"timestamp"`
	LogLevel     string `json:"log_level"`
	LogType      string `json:"log_type"`
	Event        string `json:"event"`
	AppName      string `json:"app_name"`
	Status       int    `json:"status"`
	DurationMs   string `json:"duration_ms"`
	ResponseBody string `json:"response_body"`
	RequestId    string `json:"request_id"`
}

func newResponseLog(logLevel string, logType string, responseBody string, appName string, status int, durationMs string, requestId string) *responseLog {
	currentTime := time.Now()
	return &responseLog{
		Timestamp:    currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		LogLevel:     logLevel,
		LogType:      logType,
		Event:        "response",
		AppName:      appName,
		Status:       status,
		DurationMs:   durationMs,
		ResponseBody: responseBody,
		RequestId:    requestId,
	}
}
