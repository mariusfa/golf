package tracelog

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mariusfa/golf/logging/utils"
)

var traceLogger = NewTraceLogger("APP_NAME_NOT_SET")

func NewTraceLogger(appName string) *slog.Logger {
	return utils.NewSlogger().With(
		slog.String("app_name", appName),
		slog.String("log_type", "TRACE"),
	)
}

func SetAppName(appName string) {
	traceLogger = NewTraceLogger(appName)
}

func Info(ctx context.Context, payload string) {
	username, requestId := utils.ExtractFromContext(ctx)

	traceLogger.Info(
		payload,
		slog.String("user_id", username),
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

func Error(ctx context.Context, payload string) {
	username, requestid := utils.ExtractFromContext(ctx)

	traceLogger.Error(
		payload,
		slog.String("user_id", username),
		slog.String("request_id", requestid),
	)
}

func Errorf(ctx context.Context, payload string, err error) {
	username, requestid := utils.ExtractFromContext(ctx)

	if err != nil {
		payload = fmt.Sprintf(payload+": %v", err)
	}

	traceLogger.Error(
		payload,
		slog.String("user_id", username),
		slog.String("request_id", requestid),
	)
}
