package applog

import (
	"fmt"
	"log/slog"

	"github.com/mariusfa/golf/logging/utils"
)

var appLogger = NewAppLogger("APP_NAME_NOT_SET")

func NewAppLogger(appName string) *slog.Logger {
	logger := utils.NewSlogger()
	return logger.With(
		slog.String("app_name", appName),
		slog.String("log_type", "APP"),
	)
}

func SetAppName(appName string) {
	appLogger = NewAppLogger(appName)
}

func Info(payload string) {
	appLogger.Info(payload)
}

func Infof(format string, v ...any) {
	payload := format
	if len(v) > 0 {
		payload = fmt.Sprintf(format, v...)
	}

	appLogger.Info(payload)
}

func Error(payload string) {
	appLogger.Error(payload)
}

func Errorf(payload string, error error) {
	if error != nil {
		payload = fmt.Sprintf(payload+": %v", error)
	}

	appLogger.Error(payload)
}
