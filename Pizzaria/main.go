package main

import (
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// c = context
// r = router

func main() {
	RouterSetup()
}

var pizzas []models.Pizza = []models.Pizza{}

func GetPizzas(c *gin.Context) {
	c.JSON(200, pizzas)
}

func PostPizza(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	pizzas = append(pizzas, newPizza)
}

func GetPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for _, v := range pizzas {
		if v.ID == id {
			c.JSON(200, v)
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}

func RouterSetup() {
	r := gin.Default()
	r.GET("/pizzas", GetPizzas)
	r.GET("/pizzas/:id", GetPizzasById)
	r.POST("/pizzas", PostPizza)
	r.Run()
}
