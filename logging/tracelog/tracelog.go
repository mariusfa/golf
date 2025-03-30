package tracelog

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mariusfa/golf/logging/utils"
)

var tracelogger = NewTraceLogger("APP_NAME_NOT_SET")

func NewTraceLogger(appName string) *slog.Logger {
	logger := utils.NewSlogger()
	return logger.With(
		slog.String("app_name", appName),
		slog.String("log_type", "TRACE"),
	)
}

func SetAppName(appName string) {
	tracelogger = NewTraceLogger(appName)
}

func Info(ctx context.Context, payload string) {
	username, requestId := utils.ExtractFromContext(ctx)

	tracelogger.Info(
		payload,
		slog.String("username", username),
		slog.String("request_id", requestId),
	)
}

func Infof(ctx context.Context, format string, v ...any) {
	payload := format
	if len(v) > 0 {
		payload = fmt.Sprintf(format, v...)
	}
	Info(ctx, payload)
}

func Errorf(ctx context.Context, payload string, err error) {
	username, requestid := utils.ExtractFromContext(ctx)

	tracelogger.Error(
		payload,
		slog.String("username", username),
		slog.String("request_id", requestid),
		slog.Any("error", err),
	)
}

func Error(ctx context.Context, payload string) {
	username, requestid := utils.ExtractFromContext(ctx)

	tracelogger.Error(
		payload,
		slog.String("username", username),
		slog.String("request_id", requestid),
	)
}
