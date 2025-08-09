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
	data, err := file.ReadFile(fileName)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("post", url+"/b", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	setHeaders(req)
	req.Header.Set("X-Bin-Name", name)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func UpdateBin(fileName, id string) error {
	data, err := file.ReadFile(fileName)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("put", url+"/b/"+id, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	setHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func setHeaders(req *http.Request, name string) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", config.GetConfig().Key)
	if name != "" {
		req.Header.Set("X-Bin-Name", name)
	}
}

func sendRequest(method, url, fileName, name string) {
	data, err := file.ReadFile(fileName)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	setHeaders(req, name)
}
