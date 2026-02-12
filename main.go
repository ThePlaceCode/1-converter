package main

import "fmt"

func main() {
	// Константы конвертации
	const (
		USDToEUR = 0.93
		USDToRUB = 90.5
		EURToRUB = USDToRUB / USDToEUR
	)

	fmt.Println(USDToEUR, USDToRUB, EURToRUB)
}
