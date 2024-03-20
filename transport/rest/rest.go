package rest

import (
	"github.com/gin-gonic/gin"
	"sync"
)

func StartRestService(wg *sync.WaitGroup) error {

	defer wg.Done()

	engine := gin.Default()

	// Метод обращения сheck
	engine.GET("/check", CheckRest)

	// port :8080
	err := engine.Run()

	return err

}
