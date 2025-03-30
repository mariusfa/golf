package utils

import "log/slog"

func ReplaceDefaultKeys(groups []string, attr slog.Attr) slog.Attr {
	switch attr.Key {
	case slog.TimeKey:
		return slog.Attr{Key: "timestamp", Value: attr.Value}
	case slog.LevelKey:
		return slog.Attr{Key: "log_level", Value: attr.Value}
	case slog.MessageKey:
		return slog.Attr{Key: "payload", Value: attr.Value}
	default:
		return attr
	}
}
