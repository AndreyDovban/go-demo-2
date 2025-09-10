package main

import (
	"fmt"
	"strings"
)

func main() {
	tr := make([]int, 0, 2)
	tr = append(tr, 1)
	tr = append(tr, 2)
	fmt.Println(tr)

	transaction := []int{}

	for {
		q := entryTransaction(&transaction)
		if !q {
			break
		}
	}

	fmt.Println(transaction)
	sum := getSumTransaction(&transaction)

	fmt.Println(sum)

}

func entryTransaction(trs *[]int) bool {
	var transaction int
	var questoins string

	fmt.Println("Введите трансакцию:")
	fmt.Scan(&transaction)

	*trs = append(*trs, transaction)

	fmt.Println("Продолжить?")
	fmt.Scan(&questoins)

	if strings.ToLower(questoins) == "y" || strings.ToLower(questoins) == "yes" {
		return true
	}
	return false
}

func getSumTransaction(trs *[]int) int {
	var sum int
	for _, el := range *trs {
		sum += el
	}
	return sum
}
