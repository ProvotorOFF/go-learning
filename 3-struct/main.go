package main

import (
	"3-struct/app/api"
	"3-struct/app/bins"
	"3-struct/app/storage"
	"flag"
	"fmt"
)

var binList bins.BinList
var isLoaded bool
var fileStorage *storage.Storage = storage.NewStorage()

type StorageService interface {
	SaveBins(list storage.Serializable) (bool, error)
	ReadBins() (*bins.BinList, error)
}

func main() {
	isCreate := flag.Bool("create", false, "Создает новый bin из файла")
	isUpdate := flag.Bool("update", false, "Обновляет bin из файла по id")
	isDelete := flag.Bool("delete", false, "Удаляет bin по id")
	isGet := flag.Bool("get", false, "Получение Bin по id")
	isList := flag.Bool("list", false, "Получение Bin по id")
	file := flag.String("file", "file.json", "Файл с бином")
	name := flag.String("name", "bin", "Название bin")
	id := flag.String("id", "", "Идентификатор bin")
	switch {
	case *isCreate:
		api.CreateBin(*file, *name)
	case *isUpdate:
		api.Init()
	case *isDelete:
		api.Init()
	case *isGet:
		api.Init()
	case *isList:
		api.Init()
	default:
		for {
			opt := getMenu([]string{
				"Загрузить бины из хранилища",
				"Добавить бин",
				"Сохранить бины",
				"Остановить приложение",
				"Выберите пункт меню: ",
			})
			switch opt {
			case 1:
				loadBinList(fileStorage)
			case 2:
				bin := inputBinData()
				fmt.Println("Бин успешно создан")
				if isLoaded {
					binList.Bins = append(binList.Bins, bin)
				} else {
					fmt.Println("Бинлист не прочтен, читаем автоматически")
					loadBinList(fileStorage)
					if isLoaded {
						binList.Bins = append(binList.Bins, bin)
					}
				}
			case 3:
				if isLoaded {
					fileStorage.SaveBins(&binList)
				} else {
					fmt.Println("Бинлист не загружен, нечего загружать")
				}
			case 4:
				return
			default:
				fmt.Println("Неизвестный пункт меню")
			}
		}
	}
}

func getMenu(options []string) int {
	var opt int
	for i, val := range options {
		if i < len(options)-1 {
			fmt.Printf("%d. %s\n", i+1, val)
		} else {
			fmt.Printf(" %s\n", val)
		}
	}
	fmt.Scanln(&opt)
	return opt
}

func inputBinData() bins.Bin {
	id := readString("Введите id")
	private := readBool("Введите private (true/false)")
	name := readString("Введите name")
	return bins.NewBin(id, private, name)
}

func readString(message string) string {
	var input string
	fmt.Printf("%s: ", message)
	fmt.Scan(&input)
	return input
}

func readBool(message string) bool {
	var input bool
	for {
		fmt.Printf("%s: ", message)
		_, err := fmt.Scan(&input)
		if err == nil {
			return input
		}
		fmt.Println("Ошибка ввода. Введите true или false.")
		fmt.Scan(&input)
	}
}

func loadBinList(storage StorageService) {
	storageBins, err := storage.ReadBins()
	if err != nil {
		fmt.Println("Бинлист не прочтен")
	} else {
		isLoaded = true
		binList = *storageBins
		fmt.Println("Бинлист загружен")
	}
}
