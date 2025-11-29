package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// c = context
// r = router

func main() {
	data.LoadPizzas()
	RouterSetup()
}

func GetPizzas(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(200, data.Pizzas)
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for _, v := range data.Pizzas {
		if v.ID == id {
			c.JSON(200, v)
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}

func PutPizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for i, v := range data.Pizzas {
		if v.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			data.SavePizzas()
			c.JSON(200, data.Pizzas[i])
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}

func DeletePizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for i, v := range data.Pizzas {
		if v.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizzas()
			c.JSON(200, gin.H{
				"message": "pizza deleted",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}

func PostPizza(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizzas()
	c.JSON(201, newPizza)
}

func RouterSetup() {
	r := gin.Default()

	r.GET("/pizzas/:id", GetPizzas)
	r.PUT("/pizzas/:id", PutPizza)
	r.DELETE("/pizzas/:id", DeletePizza)

	r.GET("/pizzas", GetPizzas)
	r.POST("/pizzas", PostPizza)

	r.Run()
}
