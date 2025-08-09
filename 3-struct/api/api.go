package api

import (
	"3-struct/app/config"
	"3-struct/app/file"
	"bytes"
	"fmt"
	"net/http"
)

var url string = "https://api.jsonbin.io/v3"

func Init() {
	key := config.GetConfig().Key
	fmt.Print(key)
}

func CreateBin(fileName, name string) error {
	//TODO добить бины
	// var fileStorage *storage.Storage = storage.NewStorage()
	// bin, err := file.ReadFile(fileName)
	// storageBins, err := fileStorage.ReadBins()
	// fmt.Println("Бин успешно создан")

	return sendRequest("post", url+"/b/", fileName, name)
}

func UpdateBin(fileName, id string) error {
	return sendRequest("put", url+"/b/"+id, fileName, "")
}

func DeleteBin(id string) error {
	return sendRequest("delete", url+"/b/"+id, "", "")
}

func GetBin(id string) error {
	return sendRequest("get", url+"/b/"+id, "", "")
}

func setHeaders(req *http.Request, name string) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", config.GetConfig().Key)
	if name != "" {
		req.Header.Set("X-Bin-Name", name)
	}

	return nil
}

func sendRequest(method, url, fileName, name string) error {
	var data []byte
	var err error
	var req *http.Request

	if fileName != "" {
		data, err = file.ReadFile(fileName)
		if err != nil {
			return err
		}
	}

	if data != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return err
	}

	setHeaders(req, name)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
