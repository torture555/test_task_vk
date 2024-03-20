package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"sync"
	"task/config"
	"task/transport/rest"
)

func main() {
	//Чтение прааметров
	parseParams()

	//Инициализация общего хранилища данных
	err := config.InitDB()
	if err != nil {
		slog.Error("Не удалось подключится к СУБД")
		os.Exit(1)
	}

	//Запуск REST сервиса для взаимодействия
	wg := sync.WaitGroup{}
	wg.Add(1)

	err = rest.StartRestService(&wg)
	if err != nil {
		slog.Error("Не запущен Rest сервис")
		os.Exit(1)
	}

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

// Чтение параметров приложения
func parseParams() {

	host := flag.String("host", "192.168.211.3", "Хост СУБД")
	port := flag.String("port", "5432", "Порт СУБД, по умолчанию 5432")
	DBName := flag.String("dbname", "task", "Имя БД, по умолчанию default")
	login := flag.String("login", "pavel", "Имя пользователя")
	passwd := flag.String("passwd", "pavel", "Пароль пользователя")
	typeDB := flag.String("typeDB", "postgres", "Тип используемой БД")

	timeoutCheck := flag.Int("timeout", 60, "Время проверки в секундах")
	countChecks := flag.Int("count", 10, "Количество проверок за период timeout")

	flag.Parse()

	config.Host = *host
	config.Port = *port
	config.DBName = *DBName
	config.Login = *login
	config.Passwd = *passwd
	config.TypeDB = *typeDB

	config.CheckTimeout = *timeoutCheck
	config.CheckCountBlock = *countChecks

}
