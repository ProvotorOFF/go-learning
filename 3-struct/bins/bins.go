package bins

import (
	"encoding/json"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Data      string    `json:"data"`
}

func NewBin(id string, private bool, name string, data string) Bin {
	return Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
		Data:      data,
	}
}

type BinList struct {
	Bins []Bin
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{
		Bins: bins,
	}
}

func (binList *BinList) ToBytes() ([]byte, error) {
	file, err := json.Marshal(binList)
	if err != nil {
		return nil, err
	}
	return file, nil
}
