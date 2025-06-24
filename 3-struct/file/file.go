package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckIsJSON(path string) error {
	if strings.ToLower(filepath.Ext(path)) != ".json" {
		return fmt.Errorf("not_json")
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	err := CheckIsJSON(path)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}
