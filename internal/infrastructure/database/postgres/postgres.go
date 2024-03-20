package postgres

import (
	"fmt"
	"task/config"
)

func Check(unitIndex, userID int64) (bool, error) {
	resQuery := fmt.Sprintf(CheckQuery, unitIndex, userID, config.CheckTimeout)
	raw, err := config.DB.Query(resQuery)
	if err != nil {
		return false, err
	}
	defer raw.Close()
	if !raw.Next() {
		return false, err
	}

	var res int
	err = raw.Scan(&res)
	if err != nil {
		return false, err
	}

	resBool := res > config.CheckCountBlock
	fmt.Println(res)
	return resBool, nil

}

func CreateTable(minPos, maxPos int64) error {

	resQuery := fmt.Sprintf(CreateTableQuery, minPos, minPos, minPos, maxPos, minPos, minPos)
	_, err := config.DB.Exec(resQuery)
	return err

}

func CheckHaveTable(unitIndex int64) bool {

	resQuery := fmt.Sprintf(CheckHaveTableQuery, unitIndex)
	_, err := config.DB.Query(resQuery)
	if err != nil {
		return false
	} else {
		return true
	}

}

func DeleteLastMarks(unitIndex, userID int64) error {
	resQuery := fmt.Sprintf(DeleteLastMarksQuery, unitIndex, config.DeleteTimeout, userID)
	_, err := config.DB.Exec(resQuery)
	return err
}

func InsertMarkCheck(unitIndex, userID int64) error {
	resQuery := fmt.Sprintf(InsertMarkCheckQuery, unitIndex, userID)
	_, err := config.DB.Exec(resQuery)
	return err
}
