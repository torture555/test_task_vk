package model

import (
	"context"
	"errors"
	"task/internal/usecase/check"
)

type CheckStream struct {
	UserID int64
}

func (c *CheckStream) Check(ctx context.Context, userID int64) (bool, error) {

	childCtx, cancel := context.WithCancelCause(ctx)
	c.UserID = userID

	for {
		select {
		case <-ctx.Done():
			cancel(errors.New("Завершено отменой"))
			return false, nil
		default:

			res, err := check.ReadResultCheck(c.UserID)
			if err != nil {
				cancel(errors.New("Проверка прошла неудачно"))
				return false, err
			}

			go check.InsertNewCheck(childCtx, c.UserID)
			go check.DeleteOldChecks(childCtx, c.UserID)
			return res, err
		}
	}

}
