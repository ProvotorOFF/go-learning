package api

import (
	"3-struct/app/config"
	"fmt"
)

func init() {
	key := config.GetConfig().Key
	fmt.Print(key)
}
