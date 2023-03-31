package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	type Saludo struct {
		Nombre   string `json:"nombre"`
		Apellido string `json:"apellido"`
	}

	//Router de gin con dos middlewares (logger y recovery middleware)
	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Mundo",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})

	router.POST("/saludar", func(c *gin.Context) {

		var s Saludo
		if err := c.ShouldBindJSON(&s); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		mensaje := fmt.Sprintf("Hola %v %v", s.Nombre, s.Apellido)

		c.JSON(http.StatusOK, gin.H{"mensaje": mensaje})
	})

	router.Run(":8081")
}
