package utils

import (
	"math/rand"
	"sort"
	"time"
)

// Interface e métodos para ordenar um [][]int
// Len, Less e Swap são os métodos que definem a lógica do sort
type sortItems [][]int

func (x sortItems) Len() int { return len(x) }

func (x sortItems) Less(i, j int) bool {

	var auxi float64 = float64(x[i][1]) / float64(x[i][0])

	var auxj float64 = float64(x[j][1]) / float64(x[j][0])

	return auxi < auxj
}

func (x sortItems) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Dado uma quantidade, um limite superior e um inferior gera um slice de array com valores [peso, valor] uniformes
func GetItemsWeightNValues(amount int, limitBot int, limitTop int) [][]int {
	var weight []int
	var value []int
	increase := (limitTop - limitBot) / amount

	// Preenche 2 slices com valores
	for i := limitBot; i < limitTop; i += increase {
		weight = append(weight, i)
		value = append(value, i)
	}

	// Caso os slices tenha tamanho de não o passado( um a mais ou um a menos)
	if len(weight) > amount {
		weight = weight[:len(weight)-1]
	} else if len(weight) > amount {
		weight = append(weight, limitTop)
	}
	if len(value) > amount {
		value = value[:len(value)-1]
	} else if len(value) > amount {
		value = append(value, limitTop)
	}

	// Embaralha os dois slices
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(weight), func(i, j int) { weight[i], weight[j] = weight[j], weight[i] })
	rand.Shuffle(len(value), func(i, j int) { value[i], value[j] = value[j], value[i] })

	// Cria a Slice de vetor e preenche conforme: peso[i] - valor[i]
	var items [][]int
	for _, j := range weight {
		var item [2]int
		item[0] = j
		items = append(items, item[:])
	}
	for i, j := range value {
		items[i][1] = j
	}

	// Ordena conforme valor/peso
	sort.Sort(sort.Reverse(sortItems(items)))
	return items
}

// Dado uma quantidade, um limite superior e um inferior gera um slice de array com valores [peso, valor] pseudo-aleatórios
func GetItemsWeightNValuesRandom(amount int, limitBot int, limitTop int) [][]int {
	var weight []int
	var value []int
	rand.Seed(time.Now().UnixNano())

	// Preenche 2 slices com valores
	for i := 0; i < amount; i++ {
		weight = append(weight, rand.Intn(limitTop-limitBot)+limitBot)
		value = append(value, rand.Intn(limitTop-limitBot)+limitBot)
	}

	// Cria a Slice de vetor e preenche conforme: peso[i] - valor[i]
	var items [][]int
	for _, j := range weight {
		var item [2]int
		item[0] = j
		items = append(items, item[:])
	}
	for i, j := range value {
		items[i][1] = j
	}

	// Ordena conforme valor/peso
	sort.Sort(sort.Reverse(sortItems(items)))

	return items
}

// Dado uma slice de slice gera a capacidade da mochila
func GetCapacityKnapsack(items [][]int) int {
	var capacity int
	var maxWeight int

	// Pega os pesos e vai somando na variavel capacity
	for _, item := range items {
		capacity += item[0]
		if item[0] > maxWeight {
			maxWeight = item[0]
		}
	}

	capacity /= 2

	// Caso a capacidade seja menor que o maior peso soma ela ao total
	if capacity < maxWeight {
		capacity += maxWeight
	}

	return capacity
}
