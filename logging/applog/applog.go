package applog

import (
	"fmt"
	"log/slog"

	"github.com/mariusfa/golf/logging/utils"
)

var applogger = newAppLogger("APP_NAME_NOT_SET")

func newAppLogger(appName string) *slog.Logger {
	logger := utils.NewSlogger()
	return logger.With(
		slog.String("app_name", appName),
		slog.String("log_type", "APP"),
	)
}

func SetAppName(appName string) {
	applogger = newAppLogger(appName)
}

func Info(payload string) {
	applogger.Info(payload)
}

func Infof(format string, v ...any) {
	payload := format
	if len(v) > 0 {
		payload = fmt.Sprintf(format, v...)
	}

	Info(payload)
}

func Error(payload string) {
	applogger.Error(payload)
}

func Errorf(payload string, error error) {
	applogger.Error(payload, slog.Any("error", error))
}
