package transactionlog

import (
	"context"
	"log/slog"

	"github.com/mariusfa/golf/logging/utils"
)

var transLogger = NewTransLogger("APP_NAME_NOT_SET")

func NewTransLogger(appName string) *slog.Logger {
	return utils.NewSlogger().With(
		slog.String("app_name", appName),
		slog.String("log_type", "TRANSACTION"),
	)
}

func SetAppName(appName string) {
	transLogger = NewTransLogger(appName)
}

func RequestInfo(
	ctx context.Context,
	requestMethod string,
	requestPath string,
	requestBody string,
) {
	requestId, userId := utils.ExtractFromContext(ctx)
	transLogger.Info("",
		slog.String("event", "request"),
		slog.String("request_method", requestMethod),
		slog.String("request_path", requestPath),
		slog.String("request_body", requestBody),
		slog.String("request_id", requestId),
		slog.String("usr_id", userId),
	)
}

func ResponseInfo(
	ctx context.Context,
	durationMs int,
	status int,
	responseBody string,
) {
	requestId, userId := utils.ExtractFromContext(ctx)

	transLogger.Info("",
		slog.String("event", "response"),
		slog.String("response_body", responseBody),
		slog.Int("duration_ms", durationMs),
		slog.Int("status", status),
		slog.String("request_id", requestId),
		slog.String("user_id", userId),
	)
}
