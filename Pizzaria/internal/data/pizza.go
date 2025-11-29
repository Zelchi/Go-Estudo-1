package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("Erro ao carregar as pizzas: ", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Pizzas)
	fmt.Println(Pizzas)

	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}
}

func SavePizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("Erro ao carregar as pizzas: ", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Pizzas); err != nil {
		fmt.Println("Error encodnig JSON: ", err)
		return
	}
}
