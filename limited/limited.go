package limited

import (
	"context"
	"fmt"
	"mochila/utils"
	"os"
	"strconv"
	"time"
)

// Dado uma slice com os items e a capacidade da mochila, roda testes casos que o resultado seja maior ou igual a um limitante
func Limited(items [][]int, capacity int) {

	// Com o pacote context posso contolar o timeout, então caso chegue no tempo limite, o programa encerra
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	stopPoint := false
	// Constroi o primeiro nó
	node, results, result := utils.GetInitialNode(items, capacity)

	for {
		// Caso o nó não tenha mais resultados ou o programa atinja o tempo limite, encerra o programa e escreve os valores em um log
		if stopPoint || ctx.Err() != nil {
			archive, err := os.OpenFile("logLimited.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
			if err == nil {
				archive.WriteString("Type of End of Test: " + fmt.Sprint(ctx.Err()) + "\n")
				archive.WriteString("Items: [weight, value]\n" + fmt.Sprint(items) + "\n")
				archive.WriteString("Capacity: " + strconv.FormatInt(int64(capacity), 10) + "\n")
				archive.WriteString("Result: " + strconv.FormatInt(int64(result), 10) + "\n")
				archive.WriteString("Result-Matches: " + fmt.Sprint(results) + "\n/////////////////////////////////////////////////////\n")
			}
			archive.Close()
			break
		}
		// Percorre o nó e se encontrar um valor no sentido da direita pra esquerda que seja maior que o limitante atualiza o nó
		for i := len(node) - 2; i >= 0; i-- {
			if node[i] > 0 && upperLimit(items, capacity, node, i, result) {
				node, results, result = utils.GetNewNode(items, capacity, node, i, result, results)
				break
			} else if node[i] == 0 && upperLimit(items, capacity, node, i, result) {
				stopPoint = true
			}
		}
	}

}

// Dado os items, o nó, o indice do ponto, a capacidade da mochila e o maior resultado até o momento, verifica o
// Limite superior e se ele é maior que o maior resultado dado até o momento
func upperLimit(items [][]int, capacity int, node []int, index int, resultMax int) bool {
	var result int
	var capacityAux int

	// O primeiro ponto da árvore sempre vai ser calculado
	if index == 0 {
		return true
	}

	// Somatoria valores anteriores a esse indice com o valor desse indice - 1, além de calcular a somatória para o C barra
	for i, item := range items {
		if i < index {
			result += item[1] * node[i]
			capacityAux += item[0] * node[i]
		} else if i == index {
			result += item[1] * (node[i] - 1)
			break
		}
	}
	// Calcula o C barra
	capacityAux = capacity - (items[index][0] * (node[index] - 1)) - capacityAux

	// Calcula o valor/peso para o ponto x + 1
	aux := (float64(items[index+1][1]) / float64(items[index+1][0]))

	// Calcula o limite superior
	upperlimit := (float64(result) + (aux * float64(capacityAux)))

	// verifica se o limite superior é menor ou maior que o resultado maximo atual
	if upperlimit < float64(resultMax) {
		return false
	} else {
		return true
	}
}
