package api

import (
	"3-struct/app/bins"
	"3-struct/app/config"
	"3-struct/app/file"
	"3-struct/app/storage"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var url string = "https://api.jsonbin.io/v3"

type StorageService interface {
	SaveBins(list storage.Serializable) (bool, error)
	ReadBins() (*bins.BinList, error)
}

var Errors = map[string]error{
	"storage": errors.New("STORAGE_ERROR"),
	"huynya":  errors.New("HUYNYA_ERROR"),
}

func CreateBin(fileName, name string, storageService StorageService) ([]byte, error) {
	var result struct {
		Metadata struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Private bool   `json:"private"`
		}
	}
	data, err := sendRequest("POST", url+"/b/", fileName, name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	binList, err := storageService.ReadBins()
	if err != nil {
		return nil, err
	}
	newBin := bins.NewBin(
		result.Metadata.ID, result.Metadata.Private, result.Metadata.Name,
	)
	binList.Bins = append(binList.Bins, newBin)
	_, err = storageService.SaveBins(binList)
	if err != nil {
		return nil, Errors["storage"]
	}
	return data, nil
}

func UpdateBin(fileName, id string) ([]byte, error) {
	return sendRequest("PUT", url+"/b/"+id, fileName, "")
}

func DeleteBin(id string, storageService StorageService) ([]byte, error) {
	data, err := sendRequest("DELETE", url+"/b/"+id, "", "")
	if err != nil {
		return nil, err
	}

	binList, err := storageService.ReadBins()
	if err != nil {
		return nil, err
	}

	newBinList := bins.NewBinList([]bins.Bin{})

	for _, b := range binList.Bins {
		if b.Id != id {
			newBinList.Bins = append(newBinList.Bins, b)
		}
	}

	if _, err := storageService.SaveBins(newBinList); err != nil {
		return nil, err
	}

	return data, nil
}

func GetBin(id string) ([]byte, error) {
	data, err := sendRequest("GET", url+"/b/"+id, "", "")

	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	return data, nil
}

func setHeaders(req *http.Request, name string) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", config.GetConfig().Key)
	if name != "" {
		req.Header.Set("X-Bin-Name", name)
	}

	return nil
}

func sendRequest(method, url, fileName, name string) ([]byte, error) {
	var data []byte
	var err error
	var req *http.Request

	if fileName != "" {
		data, err = file.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
	}

	if data != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	setHeaders(req, name)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respData, fmt.Errorf("ошибка запроса (%d): %s", resp.StatusCode, string(respData))
	}

	return respData, nil
}
