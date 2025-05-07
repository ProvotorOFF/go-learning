package main

import "fmt"

const USDEUR = 0.88
const USDRUB = 80.50
const EURRUB = USDRUB / USDEUR

func main() {
	base, from, to := getInput()
	calculateRate(base, from, to)
}

func calculateRate(base float64, from string, to string) float64 {
	return 1.0
}

func getInput() (float64, string, string) {
	var base float64
	var from, to string
	fmt.Scan(&base, &from, &to)
	return base, from, to
}
