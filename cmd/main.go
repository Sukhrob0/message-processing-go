package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/", func(c *gin.Context) {
        message, _ := c.GetRawData()
        fmt.Println("Получено сообщение:", string(message))
        c.JSON(http.StatusOK, gin.H{"message": "Сообщение принято"})
    })

    r.Run(":8080")
}
