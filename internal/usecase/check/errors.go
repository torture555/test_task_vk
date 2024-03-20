package check

import "log/slog"

func ReadResultCheck_Error() {
	slog.Error("Не удалось получить проверку флуда")
}

func InsertNewCheck_Error() {
	slog.Error("Не удалось добавить проверку флуда")
}

func DeleteOldChecks_Error() {
	slog.Error("Не удалось удалить старые записи проверки флуда")
}
