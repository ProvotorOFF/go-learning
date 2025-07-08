package main

import "fmt"

func main() {
	for {
		opt := getMenu([]string{
			"Загрузить бины из хранилища",
			"Добавить бин",
			"Сохранить бины",
			"Выберите пункт меню: ",
		})
		switch opt {
		case 1:
			fmt.Print(1231231231231231)
		}
	}
}

func getMenu(options []string) int {
	var opt int
	for i, val := range options {
		fmt.Printf("%d. %s", i+1, val)
		if i < len(options)-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Scanf("%d", &opt)
	return opt + 1
}
