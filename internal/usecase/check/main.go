package check

import (
	"context"
	"task/internal/infrastructure/database"
)

func ReadResultCheck(userID int64) (bool, error) {
	var res bool

	res, err := database.DB.Check(userID)
	if err != nil {
		go ReadResultCheck_Error()
	}
	return res, err
}

func InsertNewCheck(ctx context.Context, userID int64) {

	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := database.DB.InsertMarkCheck(userID)
			if err != nil {
				InsertNewCheck_Error()
			}
			return
		}
	}

}

func DeleteOldChecks(ctx context.Context, userID int64) {

	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := database.DB.DeleteLastMarks(userID)
			if err != nil {
				DeleteOldChecks_Error()
			}
			return
		}
	}

}
