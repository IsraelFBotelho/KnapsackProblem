package main

import (
	"mochila/bruteforce"
	"mochila/utils"
	"os"
	"strconv"
)

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
