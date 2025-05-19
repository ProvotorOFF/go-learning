package main

import (
	"errors"
	"fmt"
	"strings"
)

const USDEUR = 0.88
const USDRUB = 80.50
const EURRUB = USDRUB / USDEUR

func main() {
	from := getCurrency("Ввод исходной валюты", "")
	base := getBase("Ввод числа")
	to := getCurrency("Ввод целевой валюты", from)
	res, err := calculateRate(base, from, to)
	if err == nil {
		fmt.Printf("%.2f %s в %s - %.4f", base, from, to, res)
	} else {
		panic(err)
	}

}

func calculateRate(base float64, from string, to string) (float64, error) {
	switch from {
	case "USD":
		switch to {
		case "EUR":
			return base * USDEUR, nil
		case "RUB":
			return base * USDRUB, nil
		}
	case "EUR":
		switch to {
		case "USD":
			return base / USDEUR, nil
		case "RUB":
			return base * EURRUB, nil
		}
	case "RUB":
		switch to {
		case "USD":
			return base / USDRUB, nil
		case "EUR":
			return base / EURRUB, nil
		}
	}
	return 0, errors.New("INCCORRECT_CURRENCIES")
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
		currency = strings.ToUpper(currency)
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
