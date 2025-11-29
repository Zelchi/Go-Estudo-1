package main

import "fmt"

type ContaBancaria struct {
	NomeTitular     string  `json:"nome_titular"`
	NumeroAgencia   int     `json:"numero_agencia"`
	NumeroConta     string  `json:"numero_conta"`
	QuantidadeSaldo float64 `json:"quantidade_saldo"`
}

func main() {
	pessoa1 := ContaBancaria{
		NomeTitular:     "João Silva",
		NumeroAgencia:   1234,
		NumeroConta:     "56789-0",
		QuantidadeSaldo: 1500.75,
	}

	fmt.Printf("%+v\n", pessoa1)
}
