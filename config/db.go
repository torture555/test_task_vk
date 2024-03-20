package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

// Зона входных переменных
var (
	TypeDB string
	Host   string
	Port   string
	DBName string
	Login  string
	Passwd string
)

// Ограничитель юнитов по ID
const UnitSize int64 = 100000
const DeleteTimeout = 3600 // В секундах. По умолчанию 1 час

// DB - Глобальная переменная подключения
var DB *sql.DB

// Подключение к БД
func InitDB() error {

	var err error

	dataSourceName := buildDataSourceName()
	if dataSourceName == "" {
		err = errors.New("Не задан тип БД или не сформирована строка подключения")
		return err
	}

	DB, err = sql.Open(TypeDB, dataSourceName)

	return err
}

// Формирование строки подключения по типу СУБД
func buildDataSourceName() string {

	switch TypeDB {
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Host, Port, Login, Passwd, DBName)
	// Можно добавить и другие БД
	default:
		return ""
	}

}
