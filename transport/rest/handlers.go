package rest

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task/internal/model"
)

// Хендлер обработки запроса на /check
// Если не заполнено в заголовке UserID, тогда код 400
// Внутренняя ошибка 500
func CheckRest(c *gin.Context) {

	for {
		select {
		case <-c.Done():
			return

		default:
			if c.GetHeader("UserID") == "" {
				c.JSON(400, gin.H{
					"result":      false,
					"description": "Не указано поле UserID. Укажите в заголовке",
				})
				return
			}
			userIDstr := c.GetHeader("UserID")
			userID, err := strconv.ParseInt(userIDstr, 10, 64)
			if err != nil {
				c.JSON(400, gin.H{
					"result":      false,
					"description": "Не удалось преобразовать поле UserID в число",
				})
				return
			}

			stream := model.CheckStream{}
			res, err := stream.Check(c, userID)
			if err != nil {
				c.JSON(500, gin.H{
					"result":      false,
					"description": "Не удалось провести флуд контроль",
				})
				return
			}

			c.JSON(200, gin.H{
				"result": res,
			})
			return

		}
	}

}
