package file

import (
	"fmt"
	"path/filepath"
	"strings"
)

func CheckIsJSON(path string) (bool, error) {
	if strings.ToLower(filepath.Ext(path)) != ".json" {
		return false, fmt.Errorf("not_json")
	}
	return true, nil
}
