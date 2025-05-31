package main

import (
	"errors"
	"fmt"
	"strings"
)

const USDEUR = 0.88
const USDRUB = 80.50
const EURRUB = USDRUB / USDEUR

var rates = map[string]map[string]float64{
	"USD": {
		"EUR": USDEUR,
		"RUB": USDRUB,
	},
	"EUR": {
		"USD": 1 / USDEUR,
		"RUB": EURRUB,
	},
	"RUB": {
		"USD": 1 / USDRUB,
		"EUR": 1 / EURRUB,
	},
}

func main() {
	from := getCurrency("Ввод исходной валюты", "")
	base := getBase("Ввод числа")
	to := getCurrency("Ввод целевой валюты", from)
	res, err := calculateRate(base, from, to, &rates)
	if err == nil {
		fmt.Printf("%.2f %s в %s - %.4f\n", base, from, to, res)
	} else {
		panic(err)
	}

}

func calculateRate(base float64, from string, to string, rateMap *map[string]map[string]float64) (float64, error) {
	if rate, ok := (*rateMap)[from][to]; ok {
		return base * rate, nil
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
