package database

import (
	"fmt"
	"log/slog"
)

func CreateTable_Error(unitIndex int64, err error) {
	slog.Error(fmt.Sprintf("Не удалось создать таблицу %s: %s"), unitIndex, err.Error())
}

func Check_Error(userID, unitIndex int64, err error) {
	slog.Error(fmt.Sprintf("Не удалось проверить контроль флуда в таблице %s, пользователю %s: %s"), unitIndex, userID, err.Error())
}

func InsertMarkCheck_Error(userID, unitIndex int64, err error) {
	slog.Error(fmt.Sprintf("Не удалось добавить проверку флуда в таблицу %s, пользователь %s: %s"), unitIndex, userID, err.Error())
}

func DeleteLastMarks_Error(userID, unitIndex int64, err error) {
	slog.Error(fmt.Sprintf("Не удалось удалить старые записи флуда в таблице %s, пользователя %s: %s"), unitIndex, userID, err.Error())
}
