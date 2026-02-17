package main

import "fmt"

func main() {
	// Константы конвертации
	const (
		USDToEUR = 0.93
		USDToRUB = 90.5
		EURToRUB = USDToRUB / USDToEUR
	)
}

func userInput() {
	var str string
	fmt.Scan(&str)

}

func calculate(number float64, beginCurrency string, endCurrency string) float64 {
	return 0.00
}
