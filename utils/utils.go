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

// Pega os items e a capacidade e cria o primeiro nó
func GetInitialNode(items [][]int, capacity int) ([]int, [][]int, int) {
	var result int
	var results [][]int
	var node []int
	// Percorre item por item e ve quantas vezes pode ter, em sequida diminui e aumenta a capacidade atual e o resultado respectivamente
	for _, item := range items {
		times := capacity / item[0]
		capacity -= times * item[0]
		node = append(node, times)
		result += item[1] * times
	}
	// Salva e cria um slice com os resultados possiveis
	aux := make([]int, len(items))
	copy(aux, node)
	results = append(results, aux)

	return node, results, result
}

// Pega um nó ja criado, reduz um na posição passada por parametro e recalcula o novo nó
func GetNewNode(items [][]int, capacity int, node []int, index int, resultMax int, results [][]int) ([]int, [][]int, int) {
	var result int

	// Caso o índice do item seja menor que o passado, se mantém, se for igual reduz um e se for maior re-calcula
	for i, item := range items {
		if i < index {
			times := node[i]
			capacity -= item[0] * times
			node[i] = times
			result += item[1] * times
		} else if i == index {
			times := node[i] - 1
			capacity -= item[0] * times
			node[i] = times
			result += item[1] * times
		} else if i > index {
			times := capacity / item[0]
			capacity -= times * item[0]
			node[i] = times
			result += item[1] * times
		}
	}
	// Caso o resultado desse nó supere o maior resultado até o momento, sobrescreve o valor maximo e a slice de resultados
	// Caso o nó der o mesmo resultado, adiciona o nó atual no slice
	if result > resultMax {
		resultMax = result
		results = results[:0]
		aux := make([]int, len(items))
		copy(aux, node)
		results = append(results, aux)
	} else if result == resultMax {
		aux := make([]int, len(items))
		copy(aux, node)
		results = append(results, aux)
	}
	return node, results, resultMax
}
