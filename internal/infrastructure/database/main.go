package database

import (
	"database/sql"
	"errors"
	"task/config"
	"task/internal/infrastructure/database/postgres"
)

// Методы обращения к БД
type ConnectDatabase interface {
	CreateTable(lenID int64) error
	Check(userID int64) (bool, error)
	InsertMarkCheck(userID int64) error
	DeleteLastMarks(userID int64) error
}

type DataBase struct {
	*sql.DB
}

var DB = DataBase{config.DB}

func (db *DataBase) CreateTable(userID int64) error {

	minPos := GetUnitIndex(userID)
	maxPos := GetMaxPos(minPos)
	var err error

	switch config.TypeDB {
	case "postgres":
		err = postgres.CreateTable(minPos, maxPos)

	default:
		err = errors.New("Не определен тип БД")
	}

	if err != nil {
		CreateTable_Error(minPos, err)
	}
	return err

}

func (db *DataBase) Check(userID int64) (bool, error) {

	unitIndex := GetUnitIndex(userID)
	res := false
	var err error

	switch config.TypeDB {
	case "postgres":
		checkHaveTable := postgres.CheckHaveTable(unitIndex)
		if !checkHaveTable {
			err = db.CreateTable(userID)
			if err != nil {
				break
			}
			return false, nil
		}
		res, err = postgres.Check(unitIndex, userID)
	default:
		err = errors.New("Не определен тип БД")
	}

	if err != nil {
		Check_Error(userID, unitIndex, err)
	}

	return res, err
}

func (db *DataBase) InsertMarkCheck(userID int64) error {

	unitIndex := GetUnitIndex(userID)
	var err error

	switch config.TypeDB {
	case "postgres":
		err = postgres.InsertMarkCheck(unitIndex, userID)

	default:
		err = errors.New("Не определен тип БД")
	}

	if err != nil {
		InsertMarkCheck_Error(userID, unitIndex, err)
	}

	return err

}

func (db *DataBase) DeleteLastMarks(userID int64) error {

	unitIndex := GetUnitIndex(userID)
	var err error

	switch config.TypeDB {
	case "postgres":
		err = postgres.DeleteLastMarks(unitIndex, userID)

	default:
		err = errors.New("Не определен тип БД")
	}

	if err != nil {
		DeleteLastMarks_Error(userID, unitIndex, err)
	}

	return err

}

func GetUnitIndex(userID int64) int64 {
	minPos := (userID / config.UnitSize) * config.UnitSize
	if userID < config.UnitSize {
		minPos = 0
	}
	return minPos
}

func GetMaxPos(minPos int64) int64 {
	return minPos + config.UnitSize
}
