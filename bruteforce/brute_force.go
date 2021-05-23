package bruteforce

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Dado uma slice com os items e a capacidade da mochila, roda testes considadando todos os casos
func Brute_force(items [][]int, capacity int) {

	// Com o pacote context posso contolar o timeout, então caso chegue no tempo limite, o programa encerra
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Constroi o primeiro nó
	node, results, result := getInitialNode(items, capacity)

	for {
		// Caso o nó só tenha 0 ou o programa atinja o tempo limite, encerra o programa e escreve os valores em um log
		if checkEnd(node) || ctx.Err() != nil {
			archive, err := os.OpenFile("logBrute.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
			if err == nil {
				archive.WriteString("Type of End of Test: " + fmt.Sprint(err) + "\n")
				archive.WriteString("Items: [weight, value]\n" + fmt.Sprint(items) + "\n")
				archive.WriteString("Capacity: " + strconv.FormatInt(int64(capacity), 10) + "\n")
				archive.WriteString("Result: " + strconv.FormatInt(int64(result), 10) + "\n")
				archive.WriteString("Result-Matches: " + fmt.Sprint(results) + "\n/////////////////////////////////////////////////////\n")
			}
			archive.Close()
			break
		}
		// Percorre o nó e se encontrar um valor no sentido da direita pra esquerda atualiza o nó
		for i := len(node) - 2; i >= 0; i-- {
			if node[i] > 0 {
				node, results, result = getNewNode(items, capacity, node, i, result, results)
				break
			}
		}
	}

}

// Pega os items e a capacidade e cria o primeiro nó
func getInitialNode(items [][]int, capacity int) ([]int, [][]int, int) {
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

// Verifica se ainda existe um valor da lista exceto pelo ultimo
func checkEnd(node []int) bool {
	for i, element := range node {
		if element > 0 && i != (len(node)-1) {
			return false
		}
	}
	return true
}

// Pega um nó ja criado, reduz um na posição passada por parametro e recalcula o novo nó
func getNewNode(items [][]int, capacity int, node []int, index int, resultMax int, results [][]int) ([]int, [][]int, int) {
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
