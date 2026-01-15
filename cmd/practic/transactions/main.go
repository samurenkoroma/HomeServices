package main

import (
	"fmt"
	"strconv"
)

func main() {
	transactions := []float64{}
	var input string

	for {
		fmt.Print("Введите транзакцию: ")
		fmt.Scan(&input)
		if input == "q" || input == "0" {
			break
		}

		tr, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			transactions = append(transactions, tr)
		}

	}
	var result float64
	for _, value := range transactions {
		result += value
	}
	fmt.Printf("Баланс: %.2f\n", result)

	fmt.Println("Пока")
}
