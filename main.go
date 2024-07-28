package main

import "fmt"

func main() {
	// tr1 := []int{1, 2, 3}
	// tr2 := []int{4, 5, 6}
	// tr1 = append(tr1, tr2...)

	// fmt.Println(tr1)

	transactions := []float64{}

	for {

		fmt.Println("Введите значение трансакции: ")
		transaction := scanTransaction()

		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(getSum(transactions))
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Scan(&transaction)
	return transaction
}

func getSum(arr []float64) float64 {
	var sum float64
	for _, el := range arr {
		sum += el
	}
	return sum
}
