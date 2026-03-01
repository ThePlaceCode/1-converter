package main

import (
	"fmt"
)

// Константы конвертации
const (
	USDToEUR = 0.93
	USDToRUB = 90.5
	EURToRUB = USDToRUB / USDToEUR
)

func main() {
	rates := map[string]float64{
		"USDToEUR": 0.93,
		"USDToRUB": 90.5,
		"EURToRUB": USDToRUB / USDToEUR,
	}

	beginCurrency := inputBeginCurrency()
	amountCurrency := inputAmountCurrency()
	endCurrency := inputEndCurrency(beginCurrency)

	sum := calculate(rates, amountCurrency, beginCurrency, endCurrency)
	fmt.Println(sum)
}

func userInput() {
	var str string
	fmt.Scan(&str)

}

func calculate(rates map[string]float64, amount float64, beginCurrency string, endCurrency string) float64 {
	var sum float64

	rate := beginCurrency + "To" + endCurrency
	sum = amount * rates[rate]

	return sum
}

func inputBeginCurrency() string {
	var currency string

	for {
		fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
		_, err := fmt.Scan(&currency)
		if err != nil {
			fmt.Println("Такой валюты нет. Введите еще раз.")
			continue
		}

		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			return currency
		}

		fmt.Println("Такой валюты нет. Введите еще раз.")
	}
}

func inputAmountCurrency() float64 {
	var amount float64

	for {
		fmt.Print("Введите количество валюты: ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка. Введите заново.")
			continue
		}

		return amount
	}
}

func inputEndCurrency(beginCurrency string) string {
	var currency string
	var message string

	switch beginCurrency {
	case "USD":
		message = "Введите целевую валюту (EUR, RUB): "
	case "EUR":
		message = "Введите целевую валюту (USD, RUB): "
	case "RUB":
		message = "Введите целевую валюту (USD, EUR): "
	}

	for {
		fmt.Print(message)
		_, err := fmt.Scan(&currency)
		if err != nil {
			fmt.Println("Такой валюты нет. Введите еще раз.")
			continue
		}

		if currency == beginCurrency {
			fmt.Println("Целевая валюта не должна совпадать с исходной. Введите еще раз.")
			continue
		}

		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			return currency
		}

		fmt.Println("Такой валюты нет. Введите еще раз.")
	}
}
