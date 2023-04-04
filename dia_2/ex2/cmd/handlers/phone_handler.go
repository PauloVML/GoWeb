package handlers

import (
	"github.com/PauloVML/GoWeb/dia_2/ex2/internal/phone"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TelefonoHandler struct {
	//lo primero que necesito es un service
	Service phone.TelefonoService
}

func (telHandler *TelefonoHandler) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Obtener la petición/request y validarla
		var request SaveTelefonoDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			//manejo de errores
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}

		//Crear objeto
		telefono := request.toDomain()

		//el método save recibe un PUNTERO de teléfono, y un puntero hace referencia a la dirección
		//en memoria de una variable. Y para obtener la dirección en memoria de una variable
		//anteponemos el &
		if err := telHandler.Service.Save(&telefono); err != nil {
			//manejo de errores
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Producto creado con exito", "data": telefono})

	}

}

func (telHandler TelefonoHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		telefonos, _ := telHandler.Service.GetAll()
		ctx.JSON(http.StatusOK, telefonos)
	}
}
