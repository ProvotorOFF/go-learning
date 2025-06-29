package storage

import (
	"3-struct/app/bins"
	"encoding/json"
	"os"
)

func SaveBins(binList bins.BinList) (bool, error) {
	content, err := binList.ToBytes()
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

func ReadBins() (*bins.BinList, error) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}
