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
	name := flag.String("name", "binName", "Название bin")
	id := flag.String("id", "", "Идентификатор bin")

	flag.Parse()

	switch {
	case *isCreate:
		api.CreateBin(*file, *name, fileStorage)
	case *isUpdate:
		api.UpdateBin(*file, *id)
	case *isDelete:
		api.DeleteBin(*id, fileStorage)
	case *isGet:
		api.GetBin(*id)
	case *isList:
		loadBinList(fileStorage)
	default:

	}
}

func loadBinList(storage StorageService) {
	storageBins, err := storage.ReadBins()
	if err != nil {
		fmt.Println("Бинлист не прочтен")
	} else {
		isLoaded = true
		binList = *storageBins
		fmt.Println(binList)
	}
}
