package api

import (
	"3-struct/app/config"
	"3-struct/app/file"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var url string = "https://api.jsonbin.io/v3"

func CreateBin(fileName, name string) ([]byte, error) {
	var result any
	data, err := sendRequest("post", url+"/b/", fileName, name)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &result)
	fmt.Println(result)
	return data, nil
}

func UpdateBin(fileName, id string) ([]byte, error) {
	return sendRequest("put", url+"/b/"+id, fileName, "")
}

func DeleteBin(id string) ([]byte, error) {
	return sendRequest("delete", url+"/b/"+id, "", "")
}

func GetBin(id string) ([]byte, error) {
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
