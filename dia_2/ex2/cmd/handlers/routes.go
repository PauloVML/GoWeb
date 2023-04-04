package handlers

import (
	"github.com/PauloVML/GoWeb/dia_2/ex2/internal/phone"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func (router *Router) Setup() {
	//Seteamos los middlewares por default
	router.Engine.Use(gin.Logger())
	router.Engine.Use(gin.Recovery())

	//Seteo las rutas
	router.SetTelefonoRoutes()
}

func (router *Router) SetTelefonoRoutes() {

	//Setear componentesß
	repository := &phone.RepositoryImpl{}
	service := phone.TelefonoServiceImpl{Storage: repository}
	handler := TelefonoHandler{Service: &service}

	group := router.Engine.Group("/telefono")

	group.POST("/save", handler.Save())
	group.GET("", handler.GetAll())
}
