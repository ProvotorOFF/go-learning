package main

import (
	"fmt"
)

const USDEUR = 0.88
const USDRUB = 80.50
const EURRUB = USDRUB / USDEUR

func main() {
	from := getCurrency("Ввод исходной валюты", "")
	base := getBase("Ввод числа")
	to := getCurrency("Ввод целевой валюты", from)
	fmt.Printf("%.2f %s в %s - %.4f", base, from, to, calculateRate(base, from, to))
}

func calculateRate(base float64, from string, to string) float64 {
	switch from {
	case "USD":
		switch to {
		case "EUR":
			return base * USDEUR
		case "RUB":
			return base * USDRUB
		}
	case "EUR":
		switch to {
		case "USD":
			return base / USDEUR
		case "RUB":
			return base * EURRUB
		}
	case "RUB":
		switch to {
		case "USD":
			return base / USDRUB
		case "EUR":
			return base / EURRUB
		}
	}
	return 0.0
}

func getCurrency(message string, exludedCurrency string) string {
	currency := ""
	avalibleCurrencies := ""
	switch exludedCurrency {
	case "":
		avalibleCurrencies = "RUB, USD, EUR"
	case "RUB":
		avalibleCurrencies = "USD, EUR"
	case "USD":
		avalibleCurrencies = "RUB, EUR"
	case "EUR":
		avalibleCurrencies = "RUB, USD"

	}
	for {
		fmt.Printf("%s (%s): ", message, avalibleCurrencies)
		fmt.Scanln(&currency)
		switch {
		case exludedCurrency != "" && currency == exludedCurrency:
			fmt.Println("Выберите другую валюту")
			continue
		case currency == "USD":
			return currency
		case currency == "EUR":
			return currency
		case currency == "RUB":
			return currency
		}
		fmt.Println("Введите корректную валюту")
	}
}

func getBase(message string) float64 {
	base := 0.0
	for {
		fmt.Print(message, ": ")
		fmt.Scanln(&base)
		if base > 0 {
			return base
		}
		fmt.Println("Введите корректное число")
	}
}
