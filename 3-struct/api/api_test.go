package api_test

import (
	"3-struct/app/api"
	"3-struct/app/bins"
	"3-struct/app/storage"
	"errors"
	"fmt"
	"testing"
)

type FakeStorage struct {
	BinList *bins.BinList
}

func newFakeStorage() *FakeStorage {
	return &FakeStorage{
		BinList: &bins.BinList{},
	}
}

func (fs *FakeStorage) ReadBins() (*bins.BinList, error) {
	return fs.BinList, nil
}

func (fs *FakeStorage) SaveBins(list storage.Serializable) (bool, error) {
	if bl, ok := list.(*bins.BinList); ok {
		fs.BinList = bl
		return true, nil
	}
	return false, fmt.Errorf("unsupported type %T", list)
}

type BrokenStorage struct{}

func (bs *BrokenStorage) ReadBins() (*bins.BinList, error) {
	return &bins.BinList{}, nil
}
func (bs *BrokenStorage) SaveBins(list storage.Serializable) (bool, error) {
	return false, fmt.Errorf("Ошибка сохранения")
}

func TestCreateBin(t *testing.T) {
	fakeStorage := newFakeStorage()

	data, err := api.CreateBin("test.json", "testBin", fakeStorage)

	if err != nil {
		t.Errorf("Возникла ошибка %v", err)
	}

	if len(data) == 0 {
		t.Errorf("Не пришли данные")
	}

	if len(fakeStorage.BinList.Bins) != 1 {
		t.Errorf("В хранилище должен быть 1 бин, фактически - %v", len(fakeStorage.BinList.Bins))
	}

	bin := fakeStorage.BinList.Bins[0]

	if bin.Name != "testBin" {
		t.Errorf("Ожидаемое имя testBin, фактическое %v", bin.Name)
	}
}

func TestCreateBinBadStorage(t *testing.T) {
	fakeStorage := &BrokenStorage{}
	_, err := api.CreateBin("test.json", "testBin", fakeStorage)

	if !errors.Is(err, api.Errors["storage"]) {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["storage"], err)
	}
}

func TestCreateBinBadFile(t *testing.T) {
	fakeStorage := &BrokenStorage{}
	_, err := api.CreateBin("test.json", "testBin", fakeStorage)

	if !errors.Is(err, api.Errors["storage"]) {
		t.Errorf("Ожидаемая ошибка %v, фактическая - %v", api.Errors["storage"], err)
	}
}
