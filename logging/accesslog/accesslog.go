package accesslog

import (
	"context"
	"log/slog"

	"github.com/mariusfa/golf/logging/utils"
)

var accessLogger = NewAccessLogger("APP_NAME_NOT_SET")

func NewAccessLogger(appName string) *slog.Logger {
	return utils.NewSlogger().With(
		slog.String("app_name", appName),
		slog.String("log_type", "ACCESS"),
	)
}

func SetAppName(appName string) {
	accessLogger = NewAccessLogger(appName)
}

func Info(
	ctx context.Context,
	durationMs int,
	status int,
	requestPath string,
	requestMethod string,
) {
	requestId, userId := utils.ExtractFromContext(ctx)
	accessLogger.Info("",
		slog.Int("duration_ms", durationMs),
		slog.Int("status", status),
		slog.String("request_path", requestPath),
		slog.String("request_method", requestMethod),
		slog.String("request_id", requestId),
		slog.String("user_id", userId),
	)
}
