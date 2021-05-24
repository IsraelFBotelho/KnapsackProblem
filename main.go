package main

import (
	"fmt"
	"mochila/bruteforce"
	"mochila/limited"
	"mochila/utils"
	"os"
	"strconv"
)

/*
	Padrão de Run do executável:

	./main [1, 2, 3] [brute, limited, both] [limite inferior] [limite superior]

	[1, 2, 3] -> Determina se é 10, 100 ou 1000 variaveis respectivamente

	[brute, limited, both] -> Determina qual das duas versões será rodado ou se irá rodar ambas

	[limite inferior] -> Determina o menor valor das variaveis

	[limite superior] -> Determina o maior valor das variaveis
*/
func main() {
	amount_items, _ := strconv.ParseInt(os.Args[1], 0, 0)

	// Determina quantas variaveis será usado
	switch amount_items {
	case 1:
		amount_items = 10
	case 2:
		amount_items = 100
	case 3:
		amount_items = 1000
	default:
		os.Exit(-1)
	}

	limitBot, _ := strconv.ParseInt(os.Args[3], 0, 0)

	limitTop, _ := strconv.ParseInt(os.Args[4], 0, 0)

	// Constroi um slice de array com itens já ordenados
	items := utils.GetItemsWeightNValuesRandom(int(amount_items), int(limitBot), int(limitTop))

	// Retorna o valor da capacidade da mochila
	capacity := utils.GetCapacityKnapsack(items)

	// Caso queria um teste manual preencher as variaveis abaixo
	// items := [][]int{[],[],[]}
	// capacity := 0

	// Usado somente caso queira execultar um dos testes e não ambos
	switch os.Args[2] {
	case "brute":
		bruteforce.Brute_force(items, capacity)

	case "limited":
		limited.Limited(items, capacity)

	case "both":
		limited.Limited(items, capacity)
		bruteforce.Brute_force(items, capacity)

	default:
		fmt.Println("Parametro Inválido")
		os.Exit(-1)

	}

}
