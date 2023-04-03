package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {
	products := cargarData()
	sv := gin.Default()

	product := sv.Group("/products")

	product.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	product.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		var product Product
		for _, p := range products {
			if strconv.Itoa(p.Id) == id {
				product = p
				break
			}
		}
		c.JSON(http.StatusOK, product)
	})

	product.GET("/search", func(c *gin.Context) {
		priceStr := c.Query("price")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Fatal("Parametro ingresado incorrecto")
			c.JSON(http.StatusBadRequest, err.Error())
		}

		var filtredProducts []Product
		for _, p := range products {
			if p.Price < price {
				filtredProducts = append(filtredProducts, p)
			}
		}
		if len(filtredProducts) > 0 {
			c.JSON(http.StatusOK, filtredProducts)
		} else {
			c.JSON(http.StatusNotFound, nil)
		}

	})

	sv.Run()

}

func cargarData() []Product {
	var products []Product

	file, err := os.Open("./data/products.json")
	if err != nil {
		fmt.Println(errors.New("Error al abrir el archivo"))
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(errors.New("Error al leer el archivo"))
	}

	json.Unmarshal(data, &products)
	return products
}
