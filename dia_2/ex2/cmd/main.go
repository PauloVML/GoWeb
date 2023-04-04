package main

import (
	"github.com/PauloVML/GoWeb/dia_2/ex2/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()

	router := handlers.Router{
		Engine: server,
	}

	router.Setup()

	server.Run()

}
