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

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Hour)
	defer cancel()
	node, results, result := getInitialNode(items, capacity)

	for {
		if checkEnd(node) || ctx.Err() != nil {
			archive, err := os.OpenFile("logBrute.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
			if err == nil {
				archive.WriteString("Type of End of Test: " + fmt.Sprint(err) + "\n")
				archive.WriteString("Items: \n" + fmt.Sprint(items) + "\n")
				archive.WriteString("Result: " + strconv.FormatInt(int64(result), 10) + "\n")
				archive.WriteString("Result-Matches: " + fmt.Sprint(results) + "\n/////////////////////////////////////////////////////\n")
			}
			archive.Close()
			break
		}
		for i := len(node) - 2; i >= 0; i-- {
			if node[i] > 0 {
				node, results, result = getNewNode(items, capacity, node, i, result, results)
				break
			}
		}
	}

}

func getInitialNode(items [][]int, capacity int) ([]int, [][]int, int) {
	var result int
	var results [][]int
	var node []int
	for _, item := range items {
		times := capacity / item[0]
		capacity -= times * item[0]
		node = append(node, times)
		result += item[1] * times
	}
	aux := make([]int, len(items))
	copy(aux, node)
	results = append(results, aux)

	return node, results, result
}

func checkEnd(node []int) bool {
	for i, element := range node {
		if element > 0 && i != (len(node)-1) {
			return false
		}
	}
	return true
}

func getNewNode(items [][]int, capacity int, node []int, index int, resultMax int, results [][]int) ([]int, [][]int, int) {
	var result int

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

	if result > resultMax {
		resultMax = result
		results = results[:0]
		aux := make([]int, len(items))
		copy(aux, node)
		results = append(results, aux)
	} else if result == resultMax {
		results = append(results, node[:])
	}
	return node, results, resultMax
}
