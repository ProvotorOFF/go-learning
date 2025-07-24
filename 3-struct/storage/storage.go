package storage

import (
	"3-struct/app/bins"
	"encoding/json"
	"os"
)

type Serializable interface {
	ToBytes() ([]byte, error)
}

type Storage struct {
}

func NewStorage() *Storage {
	var storage Storage
	return &storage
}

func (storage *Storage) SaveBins(list Serializable) (bool, error) {
	content, err := list.ToBytes()
	if err != nil {
		return false, err
	}
	file, err := os.Create("data.json")
	if err != nil {
		return false, err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (storage *Storage) ReadBins() (*bins.BinList, error) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			return &bins.BinList{Bins: []bins.Bin{}}, nil
		}
		return nil, err
	}
	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}
