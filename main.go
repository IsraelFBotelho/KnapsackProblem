package main

import (
	"mochila/bruteforce"
	"mochila/utils"
	"os"
	"strconv"
)

/*
	Padrão de Run do executável:

	./main [1, 2, 3] [brute ou limited] [limite inferior] [limite superior]

	[1, 2, 3] -> Determina se é 10, 100 ou 1000 variaveis respectivamente

	[brute ou limited] -> Determina qual das duas versões será rodado

	[limite inferior] -> Determina o menor valor das variaveis

	[limite superior] -> Determina o maior valor das variaveis
*/
func main() {
	amount_items, _ := strconv.ParseInt(os.Args[1], 0, 0)

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

	limitBot, _ := strconv.ParseInt(os.Args[2], 0, 0)

	limitTop, _ := strconv.ParseInt(os.Args[3], 0, 0)

	items := utils.GetItemsWeightNValues(int(amount_items), int(limitBot), int(limitTop))

	capacity := utils.GetCapacityKnapsack(items)

	bruteforce.Brute_force(items, capacity)

}
