package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var operationType = map[string]func(...int){
	"SUM": func(nums ...int) {
		fmt.Println("Сумма:", sum(nums...))
	},
	"AVG": func(nums ...int) {
		fmt.Println("Среднее:", avg(nums...))
	},
	"MED": func(nums ...int) {
		fmt.Println("Медиана:", med(nums...))
	},
}

func main() {
	operation := getOperation()
	numbers, err := getNumbers()
	if err != nil {
		fmt.Printf("Ошибка валидации чисел: %s", err)
		return
	}
	operationType[operation](numbers...)
}

func getOperation() string {
	var operation string

	for {
		fmt.Print("Введите операцию (AVG, SUM, MED): ")
		fmt.Scanln(&operation)
		validated, err := validateOperation(operation)
		if err == nil {
			return validated
		}
		fmt.Println("Введена некорректная операция")
	}
}

func getNumbers() ([]int, error) {
	var input string
	var response []int

	fmt.Print("Введите масив чисел через запятую (2, 10, 5): ")
	fmt.Scanln(&input)
	numbers := strings.Split(input, ",")
	for _, number := range numbers {
		number = strings.TrimSpace(number)
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Printf("Не удалось преобразовать %s в число\n", number)
			return []int{}, err
		}
		response = append(response, num)
	}
	return response, nil
}

func validateOperation(operation string) (string, error) {
	avalibleOperations := [3]string{"AVG", "SUM", "MED"}
	for _, val := range avalibleOperations {
		if val == operation {
			return val, nil
		}
	}
	return "", errors.New("INCORRECT_OPERATION")
}

func sum(numbers ...int) int {
	total := 0
	for _, val := range numbers {
		total += val
	}
	return total
}

func avg(numbers ...int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(sum(numbers...)) / float64(len(numbers))
}

func med(numbers ...int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sort.Ints(numbers)

	return float64(numbers[len(numbers)/2])
}
