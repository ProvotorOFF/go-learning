package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

type BinList struct {
	bins []Bin
}

func NewBinList(bins []Bin) BinList {
	return BinList{
		bins: bins,
	}
}
