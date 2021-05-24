package bruteforce

import (
	"context"
	"fmt"
	"mochila/utils"
	"os"
	"strconv"
	"time"
)

// Dado uma slice com os items e a capacidade da mochila, roda testes considadando todos os casos
func Brute_force(items [][]int, capacity int) {

	startTime := time.Now()
	// Com o pacote context posso contolar o timeout, então caso chegue no tempo limite, o programa encerra
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Hour)
	defer cancel()
	// Constroi o primeiro nó
	node, results, result := utils.GetInitialNode(items, capacity)

	for {
		// Caso o nó só tenha 0 ou o programa atinja o tempo limite, encerra o programa e escreve os valores em um log
		if checkEnd(node) || ctx.Err() != nil {
			archive, err := os.OpenFile("logBrute.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
			if err == nil {
				archive.WriteString("Type of End of Test: " + fmt.Sprint(ctx.Err()) + "\n")
				archive.WriteString("Total Execution Time: " + time.Since(startTime).String() + "\n")
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
				node, results, result = utils.GetNewNode(items, capacity, node, i, result, results)
				break
			}
		}
	}

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
