package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code_value" binding:"required"`
	IsPublished bool    `json:"is_published,omitempty"`
	Expiration  string  `json:"expiration" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func main() {

	var products = []Product{}
	var id = 1

	sv := gin.Default()

	prods := sv.Group("/products")

	prods.POST("/add", func(c *gin.Context) {

		var product Product

		//valido que estén todos los campos
		if err := c.ShouldBind(&product); err != nil {
			c.JSON(http.StatusBadRequest, "Complete todos los campos")
			return
		}

		//valido código único
		for _, p := range products {
			if p.CodeValue == product.CodeValue {
				c.JSON(http.StatusConflict, "El CodeValue debe ser único")
				return
			}
		}

		//valido fecha
		if validDate := validarFecha(product.Expiration); validDate == false {
			c.JSON(http.StatusBadRequest, "Debe ingresar una fecha válida")
		}

		product.Id = id
		id++

		products = append(products, product)

		c.JSON(http.StatusOK, product)

	})

	prods.GET("/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		var product Product
		var found bool = false

		if err != nil {
			c.JSON(http.StatusBadRequest, "Ingrese un id válido")
			return
		}

		for _, p := range products {
			if p.Id == id {
				product = p
				found = true
			}
		}

		if found == true {
			c.JSON(http.StatusOK, product)
		} else {
			c.JSON(http.StatusNotFound, "No se encontró un producto con ese ID")
		}

	})

	sv.Run()

}

func validarFecha(fecha string) bool {
	partes := strings.Split(fecha, "/")
	if len(partes) != 3 {
		return false
	}
	dia, err := strconv.Atoi(partes[0])
	if err != nil || dia < 1 || dia > 31 {
		return false
	}
	mes, err := strconv.Atoi(partes[1])
	if err != nil || mes < 1 || mes > 12 {
		return false
	}
	año, err := strconv.Atoi(partes[2])
	if err != nil && año < 0 && año < 3000 {
		return false
	}
	return true
}
